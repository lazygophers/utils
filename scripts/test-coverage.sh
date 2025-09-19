#!/bin/bash

# Test Coverage Script
# This script runs tests and generates coverage reports locally

set -e

echo "🧪 Running test coverage..."

# Ensure we're in the project root
cd "$(dirname "$0")/.."

# Run tests with coverage
echo "Running tests with coverage..."
make test-coverage

# Get coverage percentage
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}')
echo "📊 Total Coverage: ${COVERAGE}%"

# Check if coverage meets threshold
THRESHOLD=70
if awk "BEGIN {exit !($COVERAGE < $THRESHOLD)}"; then
    echo "❌ Coverage ${COVERAGE}% is below threshold ${THRESHOLD}%"
    exit 1
else
    echo "✅ Coverage ${COVERAGE}% meets threshold ${THRESHOLD}%"
fi

# Generate HTML report
echo "📄 Generating HTML coverage report..."
go tool cover -html=coverage.out -o coverage.html
echo "📄 Coverage report saved to coverage.html"

# Generate badge URL
if awk "BEGIN {exit !($COVERAGE >= 90)}"; then
    COLOR="brightgreen"
elif awk "BEGIN {exit !($COVERAGE >= 80)}"; then
    COLOR="green"
elif awk "BEGIN {exit !($COVERAGE >= 70)}"; then
    COLOR="yellow"
elif awk "BEGIN {exit !($COVERAGE >= 60)}"; then
    COLOR="orange"
else
    COLOR="red"
fi

BADGE_URL="https://img.shields.io/badge/coverage-${COVERAGE}%25-${COLOR}"
echo "🏷️  Badge URL: ${BADGE_URL}"

echo "🎉 Coverage check completed successfully!"