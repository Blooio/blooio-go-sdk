# Changelog

## 0.4.0 (2026-09-03)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/Blooio/blooio-go-sdk/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([234527a](https://github.com/Blooio/blooio-go-sdk/commit/234527a667a47f879454a4c5ba32772c50ffa639))
* **api:** api update ([a42b85a](https://github.com/Blooio/blooio-go-sdk/commit/a42b85aaa4673d33769ee0d95625ce049398698a))
* **api:** api update ([6459e15](https://github.com/Blooio/blooio-go-sdk/commit/6459e15739e3d5a5ba27076ff9cc27644b4dd866))
* **api:** api update ([203ff5b](https://github.com/Blooio/blooio-go-sdk/commit/203ff5b516667c165028419061656e6acfb23522))
* **api:** api update ([395345e](https://github.com/Blooio/blooio-go-sdk/commit/395345e84c716496bf7aef136eaef8ea14a6608b))
* **api:** api update ([ec41f65](https://github.com/Blooio/blooio-go-sdk/commit/ec41f651c092001361af990de10687b9225e7e4c))
* **api:** api update ([ee4eb59](https://github.com/Blooio/blooio-go-sdk/commit/ee4eb59adf0eb657a7590b12b4519d1dff82b8f1))
* **api:** api update ([e666c6d](https://github.com/Blooio/blooio-go-sdk/commit/e666c6da0e0d607598428eee9258bcae3337158f))
* **api:** api update ([f790ad0](https://github.com/Blooio/blooio-go-sdk/commit/f790ad0d20012104e653711b715e4d212f721ab0))
* **api:** api update ([70f1ebc](https://github.com/Blooio/blooio-go-sdk/commit/70f1ebc6ddaa26ac1f869e3524857272f7bfaf59))
* **api:** api update ([5903e4c](https://github.com/Blooio/blooio-go-sdk/commit/5903e4c7e8b528004c12137dd00915ff09bb7a82))
* **api:** api update ([7b3ab84](https://github.com/Blooio/blooio-go-sdk/commit/7b3ab8467bcaac1dfcd8472959b6b7ca220a9cc7))
* **api:** api update ([da8f72f](https://github.com/Blooio/blooio-go-sdk/commit/da8f72f0d74b75e5b5a3f25671f4d6dedcb4425d))
* **api:** api update ([61ca1e1](https://github.com/Blooio/blooio-go-sdk/commit/61ca1e12b1c8c392daff60e0a542c47d9e844a55))
* **api:** api update ([ba4c87c](https://github.com/Blooio/blooio-go-sdk/commit/ba4c87cb8d30c9e7ae85964de7b2ad0ce84e370a))
* **api:** api update ([705c80e](https://github.com/Blooio/blooio-go-sdk/commit/705c80e29120c8958043f42f7ae40592b555c3ad))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([4000ed7](https://github.com/Blooio/blooio-go-sdk/commit/4000ed7222957882f041449f21ba4c65d2f002fe))

## 0.3.0 (2026-05-14)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/Blooio/blooio-go-sdk/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([f10e230](https://github.com/Blooio/blooio-go-sdk/commit/f10e230c39840235f69884e3df177fe2e7e99b2a))
* **api:** api update ([78cdf6d](https://github.com/Blooio/blooio-go-sdk/commit/78cdf6d936de0de54ab278392a677630e6678700))
* **api:** api update ([5182b9f](https://github.com/Blooio/blooio-go-sdk/commit/5182b9fe93ca919f3e7814fce47a8dab1e3bc5df))
* **api:** manual updates ([fb71e34](https://github.com/Blooio/blooio-go-sdk/commit/fb71e3471f89f67ea5adf34e7030eb6e8068f79a))
* **client:** optimize json encoder for internal types ([0133c87](https://github.com/Blooio/blooio-go-sdk/commit/0133c873e476901d069a2b31e4ce82832c07e39f))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([5a00ed6](https://github.com/Blooio/blooio-go-sdk/commit/5a00ed6526303c7c30fd28767a944a521393cafb))


### Chores

* avoid embedding reflect.Type for dead code elimination ([9efd409](https://github.com/Blooio/blooio-go-sdk/commit/9efd40915daed63a48e150b964e82404e0094d76))
* redact api-key headers in debug logs ([ced785c](https://github.com/Blooio/blooio-go-sdk/commit/ced785ca7d34519918cbdd1ca2983b0e54e9413b))

## 0.2.0 (2026-04-30)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/Blooio/blooio-go-sdk/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([fcb3697](https://github.com/Blooio/blooio-go-sdk/commit/fcb3697d62445486b1a68cd99c500e11b23ee888))
* **client:** add a convenient param.SetJSON helper ([c4a6d48](https://github.com/Blooio/blooio-go-sdk/commit/c4a6d48d6f5d34d2fc84c71163b761b4c7eb7510))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([6831a45](https://github.com/Blooio/blooio-go-sdk/commit/6831a4569ba852783caeefe5deb9a8e03b71ff86))
* **docs:** add missing pointer prefix to api.md return types ([95acddd](https://github.com/Blooio/blooio-go-sdk/commit/95acddd02a22f4d6ed31370c21d12c69aebe08c2))
* **encoder:** correctly serialize NullStruct ([efc8462](https://github.com/Blooio/blooio-go-sdk/commit/efc8462d3441507831f20c5815f568183541a265))


### Chores

* **internal:** codegen related update ([d768da4](https://github.com/Blooio/blooio-go-sdk/commit/d768da4794ff5a450c60f1e2b46536a50c5a8267))
* **internal:** codegen related update ([663d0b3](https://github.com/Blooio/blooio-go-sdk/commit/663d0b338798b91e89b0ea743e2f3f317fc5a079))
* **internal:** codegen related update ([2a90479](https://github.com/Blooio/blooio-go-sdk/commit/2a90479986543701a167a5e77d3e41890436623d))
* **internal:** codegen related update ([c8452f0](https://github.com/Blooio/blooio-go-sdk/commit/c8452f0159f5d6869a17826b6c3c34760986f6f6))
* **internal:** remove mock server code ([8996b9b](https://github.com/Blooio/blooio-go-sdk/commit/8996b9bb3d081e4b2618e23195c30f6d3d6a1e8d))
* **internal:** update `actions/checkout` version ([3bf52b4](https://github.com/Blooio/blooio-go-sdk/commit/3bf52b490fd6423944e9692a456ad25c80ff3b55))
* update mock server docs ([9c8920d](https://github.com/Blooio/blooio-go-sdk/commit/9c8920d70cfe03fe5eea35ecc7eb157fab9103a4))
* update SDK settings ([4ed52c1](https://github.com/Blooio/blooio-go-sdk/commit/4ed52c1565191759d64cde8da3dbfbc094027a2c))

## 0.1.0 (2025-12-19)

Full Changelog: [v0.0.4...v0.1.0](https://github.com/Blooio/blooio-go-sdk/compare/v0.0.4...v0.1.0)

### Features

* **encoder:** support bracket encoding form-data object members ([99b95f9](https://github.com/Blooio/blooio-go-sdk/commit/99b95f9e9b25b5a09e1c0f54625cc723ebe2e870))


### Bug Fixes

* skip usage tests that don't work with Prism ([8d1fcec](https://github.com/Blooio/blooio-go-sdk/commit/8d1fcec85828458db33fe9498fa8af99d448e3cf))


### Chores

* add float64 to valid types for RegisterFieldValidator ([f4807d4](https://github.com/Blooio/blooio-go-sdk/commit/f4807d423c9ecd5ad7053547e8d789afc7f3bef8))

## 0.0.4 (2025-12-06)

Full Changelog: [v0.0.3...v0.0.4](https://github.com/Blooio/blooio-go-sdk/compare/v0.0.3...v0.0.4)

### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([85fa44a](https://github.com/Blooio/blooio-go-sdk/commit/85fa44aef750f3748a622e1958c21e19c7302cc6))
* **mcp:** correct code tool API endpoint ([29e2eac](https://github.com/Blooio/blooio-go-sdk/commit/29e2eac46bdde7b398adad4bf8ad90652080595b))
* rename param to avoid collision ([6c80a83](https://github.com/Blooio/blooio-go-sdk/commit/6c80a83a04a5b130ad9f4d1754a7e6425f85bb38))


### Chores

* elide duplicate aliases ([b4e3d93](https://github.com/Blooio/blooio-go-sdk/commit/b4e3d93de1edb856639d4c8d95cdaafd4eb39b38))
* fix empty interfaces ([ee13cb1](https://github.com/Blooio/blooio-go-sdk/commit/ee13cb13cc0b63fdeef92afb2fd5fe8ed3693270))
* **internal:** codegen related update ([aacbb91](https://github.com/Blooio/blooio-go-sdk/commit/aacbb916c9e53547a0cc963869eafde646fc6429))

## 0.0.3 (2025-11-12)

Full Changelog: [v0.0.2...v0.0.3](https://github.com/Blooio/blooio-go-sdk/compare/v0.0.2...v0.0.3)

### Chores

* bump gjson version ([93247c2](https://github.com/Blooio/blooio-go-sdk/commit/93247c273b7598bbde55ca7af3bbe98fcfe4ab30))
* **internal:** grammar fix (it's -&gt; its) ([ba6f605](https://github.com/Blooio/blooio-go-sdk/commit/ba6f605039a83dd1d1b5b5d5e1378a5a70f4304b))

## 0.0.2 (2025-10-17)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/Blooio/blooio-go-sdk/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([b3de0ef](https://github.com/Blooio/blooio-go-sdk/commit/b3de0ef8c25a4b0f36f10ed736c7d1ced7a718e4))
* update SDK settings ([113c18d](https://github.com/Blooio/blooio-go-sdk/commit/113c18d5484bd8ac131f85553c94e3f8c13f5705))
