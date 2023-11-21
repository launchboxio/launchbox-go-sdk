# Changelog

All notable changes to this project will be documented in this file.

## [1.5.0](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.4.1...v1.5.0) (2023-11-21)


### Features

* Add endpoints for vcs_connections ([092bcb6](https://github.com/launchboxio/launchbox-go-sdk/commit/092bcb61a5fc617784f499b42465cda731385b25))
* Add service endpoints ([9193d26](https://github.com/launchboxio/launchbox-go-sdk/commit/9193d267a0fd1f1aa6bf75fe632fbf2a0e713c36))

### [1.4.1](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.4.0...v1.4.1) (2023-11-15)


### Bug Fixes

* Add version fields to addon response ([30b27f0](https://github.com/launchboxio/launchbox-go-sdk/commit/30b27f0dfaaa0353707108fa20ea9f8d71fff4fe))

## [1.4.0](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.3.6...v1.4.0) (2023-11-15)


### Features

* Add GetProjectManifest ([051523c](https://github.com/launchboxio/launchbox-go-sdk/commit/051523c9a231dade3040db607d3cde1b539989bf))


### Bug Fixes

* Add user object as well ([736682c](https://github.com/launchboxio/launchbox-go-sdk/commit/736682cec858fd617b9e5e6cafe3aa9dc0ed5326))

### [1.3.6](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.3.5...v1.3.6) (2023-11-15)


### Bug Fixes

* Add kubernetes_version to both projects.create and projects.update ([00eb57a](https://github.com/launchboxio/launchbox-go-sdk/commit/00eb57ad19511fe85a83402108a16734a9d2edda))

### [1.3.5](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.3.4...v1.3.5) (2023-11-15)


### Bug Fixes

* Add KubernetesVersion to Project resource ([65e222b](https://github.com/launchboxio/launchbox-go-sdk/commit/65e222bc66188069c21919c9936310603d41eb45))

### [1.3.4](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.3.3...v1.3.4) (2023-11-14)


### Bug Fixes

* Flesh our clusters.update and clusters.delete ([c5a66e0](https://github.com/launchboxio/launchbox-go-sdk/commit/c5a66e0f6404b2e3e426314ecc3e050c981cbb4c))

### [1.3.3](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.3.2...v1.3.3) (2023-11-14)


### Bug Fixes

* Allow setting of cluster ID when creating a project ([d1f62b9](https://github.com/launchboxio/launchbox-go-sdk/commit/d1f62b98a35f22fad07c2c1c806ede856057f3cd))

### [1.3.2](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.3.1...v1.3.2) (2023-11-14)


### Bug Fixes

* Add domain to clusters.update ([2858085](https://github.com/launchboxio/launchbox-go-sdk/commit/2858085b726deb1dbe80f6b0f33cbd609434b0ce))

### [1.3.1](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.3.0...v1.3.1) (2023-11-14)


### Bug Fixes

* Add status field to clusters.update ([8eb499c](https://github.com/launchboxio/launchbox-go-sdk/commit/8eb499c34a27279c6e6c20aac1b95895efb19d4e))

## [1.3.0](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.2.0...v1.3.0) (2023-11-13)


### Features

* SDK functionality for ProjectAddon management ([6589d41](https://github.com/launchboxio/launchbox-go-sdk/commit/6589d41176a6947fc6d7f400f84c82b1da5ef12b))

## [1.2.0](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.1.2...v1.2.0) (2023-11-08)


### Features

* **clusters:** When viewing clusters, return client credentials if populated ([0004c27](https://github.com/launchboxio/launchbox-go-sdk/commit/0004c27ff9e58819d70fdab107f8293f2308d0f6))

### [1.1.2](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.1.1...v1.1.2) (2023-11-07)


### Bug Fixes

* **projects:** Support updating of projects ([bf37123](https://github.com/launchboxio/launchbox-go-sdk/commit/bf37123e3dcd4d32a823a406bd38b3922b5ed612))

### [1.1.1](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.1.0...v1.1.1) (2023-11-07)


### Bug Fixes

* **clusters:** Use PATCH /clusters instead of POST /ping ([56b1f50](https://github.com/launchboxio/launchbox-go-sdk/commit/56b1f507ca5b824b2941f35eaf1b74a3416c7509))

## [1.1.0](https://github.com/launchboxio/launchbox-go-sdk/compare/v1.0.0...v1.1.0) (2023-11-07)


### Features

* **auth:** Fix provider chain to support client credentials ([4776bd2](https://github.com/launchboxio/launchbox-go-sdk/commit/4776bd292f5f7b09ab39a8743b5a8918597c280e))

## 1.0.0 (2023-11-07)


### Bug Fixes

* **ci:** Add semantic-release ([beb4df7](https://github.com/launchboxio/launchbox-go-sdk/commit/beb4df7a4f179442a6cbc099e6c0f93770882f13))
* **ci:** Adjust file paths for release ([8a6007a](https://github.com/launchboxio/launchbox-go-sdk/commit/8a6007a49a3825be9f8fae2283fbc87e8713e170))
* **ci:** Fix releaserc.json syntax ([0ce9f9a](https://github.com/launchboxio/launchbox-go-sdk/commit/0ce9f9ac3fc5e696ca596eed265b174831ca581c))
* **ci:** Trigger releases on releaserc.json changes as well ([5b2df63](https://github.com/launchboxio/launchbox-go-sdk/commit/5b2df63ffe747d3256796c7b98d50e72d450bca5))
