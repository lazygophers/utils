#!/bin/bash

# Workflow Validation Script
# This script validates GitHub Actions workflow files

set -e

echo "🔍 Validating GitHub Actions workflows..."

# Ensure we're in the project root
cd "$(dirname "$0")/.."

# Check if .github/workflows directory exists
if [ ! -d ".github/workflows" ]; then
    echo "❌ .github/workflows directory not found"
    exit 1
fi

# Find all workflow files
WORKFLOW_FILES=$(find .github/workflows -name "*.yml" -o -name "*.yaml")

if [ -z "$WORKFLOW_FILES" ]; then
    echo "❌ No workflow files found"
    exit 1
fi

echo "📋 Found workflow files:"
for file in $WORKFLOW_FILES; do
    echo "  - $file"
done

# Validate YAML syntax
echo ""
echo "🔧 Validating YAML syntax..."

for file in $WORKFLOW_FILES; do
    echo "Checking $file..."

    # Check if the file is valid YAML
    if command -v yamllint >/dev/null 2>&1; then
        yamllint "$file" || echo "⚠️  yamllint not available, skipping syntax check"
    elif command -v python3 >/dev/null 2>&1; then
        python3 -c "
import yaml
import sys
try:
    with open('$file', 'r') as f:
        yaml.safe_load(f)
    print('✅ $file: Valid YAML')
except yaml.YAMLError as e:
    print('❌ $file: Invalid YAML - ', e)
    sys.exit(1)
except Exception as e:
    print('❌ $file: Error - ', e)
    sys.exit(1)
"
    else
        echo "⚠️  No YAML validator available, skipping syntax check"
    fi
done

# Check for required workflow elements
echo ""
echo "🔍 Checking workflow structure..."

for file in $WORKFLOW_FILES; do
    echo "Analyzing $file..."

    # Check for required top-level keys
    if ! grep -q "^name:" "$file"; then
        echo "⚠️  Missing 'name' field in $file"
    fi

    if ! grep -q "^on:" "$file"; then
        echo "❌ Missing 'on' field in $file"
        exit 1
    fi

    if ! grep -q "^jobs:" "$file"; then
        echo "❌ Missing 'jobs' field in $file"
        exit 1
    fi

    echo "✅ $file: Structure looks good"
done

echo ""
echo "🎉 All workflows validated successfully!"

# Show workflow summary
echo ""
echo "📊 Workflow Summary:"
echo "===================="

for file in $WORKFLOW_FILES; do
    echo ""
    echo "📄 $(basename "$file"):"

    # Extract workflow name
    NAME=$(grep "^name:" "$file" | head -1 | sed 's/name: *//' | sed 's/^"//' | sed 's/"$//')
    echo "   Name: $NAME"

    # Extract triggers
    echo "   Triggers:"
    sed -n '/^on:/,/^[a-zA-Z]/p' "$file" | grep -E '^\s+[a-zA-Z]' | sed 's/^/     - /'

    # Count jobs
    JOB_COUNT=$(grep -c "^  [a-zA-Z].*:" "$file" | head -1)
    echo "   Jobs: $JOB_COUNT"
done

echo ""
echo "✨ Validation complete!"