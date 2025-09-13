# Security Policy

## Supported Versions

We actively maintain and provide security updates for the following versions of LazyGophers Utils:

| Version | Supported          | Go Version Required | Status           |
| ------- | ------------------ | ------------------- | ---------------- |
| 1.x.x   | :white_check_mark: | Go 1.24+           | Active Development |
| 0.x.x   | :white_check_mark: | Go 1.24+           | Security Fixes Only |

## Security Considerations

### Cryptographic Components

This library includes cryptographic utilities in the `cryptox` package:

- **AES encryption/decryption**: Uses standard Go crypto library
- **RSA key generation and operations**: Implements industry-standard practices
- **Hash functions**: SHA-256, SHA-512, and other secure algorithms
- **Blowfish and ChaCha20**: Additional encryption algorithms
- **PGP/GPG**: OpenPGP implementation for secure messaging

⚠️ **Important**: While we follow security best practices, always conduct your own security review before using cryptographic functions in production environments.

### Input Validation

The utility functions in this library (especially in `candy`, `stringx`, and `config` packages) perform type conversions and data parsing. While we implement defensive programming practices:

- Input sanitization is performed where applicable
- Type conversion failures are handled gracefully
- Configuration loading includes validation steps

### Dependencies

We regularly audit our dependencies for known vulnerabilities:

- All dependencies are pinned to specific versions
- We use `go mod tidy` and `go mod verify` in our CI/CD pipeline
- Security scanning is performed via golangci-lint and gosec

## Reporting a Vulnerability

### How to Report

If you discover a security vulnerability in LazyGophers Utils, please report it responsibly:

1. **Email**: Send details to `security@lazygophers.com` (if available) or create a private issue
2. **GitHub Security Advisory**: Use GitHub's private vulnerability reporting feature
3. **Direct Contact**: Contact the maintainers directly through GitHub

### Information to Include

Please include the following information in your report:

- **Description**: Clear description of the vulnerability
- **Location**: Specific package/file/function affected
- **Impact**: Potential security impact and attack vectors
- **Reproduction**: Steps to reproduce the vulnerability
- **Suggested Fix**: If you have ideas for remediation

### Response Timeline

We are committed to addressing security issues promptly:

- **Acknowledgment**: Within 48 hours of report
- **Initial Assessment**: Within 5 business days
- **Resolution Timeline**: 
  - Critical vulnerabilities: 7-14 days
  - High severity: 14-30 days
  - Medium/Low severity: 30-90 days

### Disclosure Policy

We follow responsible disclosure practices:

1. We will work with you to understand and reproduce the issue
2. We will develop and test a fix
3. We will coordinate disclosure timing with you
4. Credit will be given to researchers who report vulnerabilities responsibly

## Security Best Practices for Users

### General Guidelines

- **Keep Updated**: Always use the latest version of the library
- **Review Code**: Conduct security reviews for production use cases
- **Validate Inputs**: Always validate external inputs in your applications
- **Follow Principles**: Apply defense-in-depth security principles

### Cryptographic Usage

When using the `cryptox` package:

- **Key Management**: Use secure key generation and storage practices
- **Random Numbers**: Ensure proper entropy for cryptographic operations
- **Algorithm Selection**: Choose appropriate algorithms for your threat model
- **Implementation**: Follow cryptographic best practices in your application

### Configuration Security

When using the `config` package:

- **File Permissions**: Restrict access to configuration files containing secrets
- **Environment Variables**: Use secure methods for storing sensitive configuration
- **Validation**: Always validate configuration data before use

## Security Testing

### Automated Security Scanning

Our CI/CD pipeline includes:

- **Static Analysis**: golangci-lint with security-focused linters
- **Vulnerability Scanning**: gosec for Go-specific security issues
- **Dependency Scanning**: Regular checks for vulnerable dependencies
- **Code Quality**: Comprehensive linting and testing

### Manual Security Review

We perform manual security reviews for:

- All cryptographic implementations
- Input validation and parsing logic
- Error handling and information disclosure
- Authentication and authorization patterns

## Contact Information

For security-related questions or concerns:

- **Project Repository**: [https://github.com/lazygophers/utils](https://github.com/lazygophers/utils)
- **Issues**: Use GitHub Issues for non-security bugs
- **Security Reports**: Follow the vulnerability reporting process above

## Changelog

### Security Updates

We maintain a security changelog for transparency:

- **Version 1.0.x**: Initial security review and hardening
- **Future versions**: Security updates will be documented here

---

**Note**: This security policy is subject to updates. Please check the latest version in the repository.

Last updated: 2025-09-13