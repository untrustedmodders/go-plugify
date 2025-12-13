# Changelog

## [1.4.4](https://github.com/untrustedmodders/go-plugify/compare/v1.4.3...v1.4.4) (2025-12-13)


### Bug Fixes

* changes in api ([ba850eb](https://github.com/untrustedmodders/go-plugify/commit/ba850eb834d6f0dda26a4d62d291705ceaaa7b2e))
* remove get method ptr ([b8d60c8](https://github.com/untrustedmodders/go-plugify/commit/b8d60c8da63a955262ae3f8061c8b18df5a8488e))

## [1.4.3](https://github.com/untrustedmodders/go-plugify/compare/v1.4.2...v1.4.3) (2025-12-08)


### Bug Fixes

* leaking memory on reassigning existing variant ([7e31f2e](https://github.com/untrustedmodders/go-plugify/commit/7e31f2e282d57922b31d3b9307bf1059915080a1))

## [1.4.2](https://github.com/untrustedmodders/go-plugify/compare/v1.4.1...v1.4.2) (2025-12-07)


### Bug Fixes

* change release type to go ([7685b82](https://github.com/untrustedmodders/go-plugify/commit/7685b829f74f29a7f76772930bb197412f914619))

## [1.4.1](https://github.com/untrustedmodders/go-plugify/compare/v1.4.0...v1.4.1) (2025-12-06)


### Bug Fixes

* remove unused functions ([f142917](https://github.com/untrustedmodders/go-plugify/commit/f142917e65d54fcbc629ac6abcc5ed0aaf8ebb79))

## [1.4.0](https://github.com/untrustedmodders/go-plugify/compare/v1.3.3...v1.4.0) (2025-12-06)


### Features

* change api for string getters ([069f89e](https://github.com/untrustedmodders/go-plugify/commit/069f89e3815f4d970441e4d0bee3de7e403624e5))


### Bug Fixes

* remove matrix4x4 from any as it not initially was supported ([7d8d362](https://github.com/untrustedmodders/go-plugify/commit/7d8d362e397381b1469cfaba6de81d440fc21545))

## [1.3.3](https://github.com/untrustedmodders/go-plugify/compare/v1.3.2...v1.3.3) (2025-11-29)


### Bug Fixes

* add dependencies and conflicts to manifest generator ([09511bb](https://github.com/untrustedmodders/go-plugify/commit/09511bbf71cd58db4b81dc9427f40100fb0f6bfe))

## [1.3.2](https://github.com/untrustedmodders/go-plugify/compare/v1.3.1...v1.3.2) (2025-11-29)


### Bug Fixes

* bump golang version ([3f5b32b](https://github.com/untrustedmodders/go-plugify/commit/3f5b32b0fe22f68f5ffe2312218a7a779e6a2586))
* **claude:** add [@brief](https://github.com/brief) for main doc ([73f6d4c](https://github.com/untrustedmodders/go-plugify/commit/73f6d4c0a5db3bcfee658ab46f0968bba4db2d70))
* **claude:** add doxygen parser ([6a5d9db](https://github.com/untrustedmodders/go-plugify/commit/6a5d9db22fa8ea956fec7b72472e66bc62c654fe))
* **claude:** add extraction of documentation for enums and generate ([85ce6cd](https://github.com/untrustedmodders/go-plugify/commit/85ce6cdbd6eef7e0da632c594cce11c7ab45edfc))
* **claude:** delegate documentation parsing ([eeefb7f](https://github.com/untrustedmodders/go-plugify/commit/eeefb7fabda8033937e0a773c9bc794a399cb1d7))
* **claude:** enum description generation ([43a2633](https://github.com/untrustedmodders/go-plugify/commit/43a2633d87888447cce38db73c8a4dfde91a8591))
* **claude:** enum description generation (2) ([7357b87](https://github.com/untrustedmodders/go-plugify/commit/7357b873e8b051eb2030c205c0bb9efe798053f7))
* **claude:** warnings on unused ([8428136](https://github.com/untrustedmodders/go-plugify/commit/8428136eed6faa0a849908375e58e752a6a592b2))
* matrix4x4 formatting ([f8decc2](https://github.com/untrustedmodders/go-plugify/commit/f8decc21953038a26bd3b3dbc05356fd43a6c245))

## [1.3.1](https://github.com/untrustedmodders/go-plugify/compare/v1.3.0...v1.3.1) (2025-11-17)


### Bug Fixes

* **claude:** fix issues with generator.go ([387517c](https://github.com/untrustedmodders/go-plugify/commit/387517cb5ac615838d6690d67584a5d947fce738))
* some gen improvements ([f5b0a2c](https://github.com/untrustedmodders/go-plugify/commit/f5b0a2c7c27f4af20caf13d502c28624dd565651))

## [1.3.0](https://github.com/untrustedmodders/go-plugify/compare/v1.2.4...v1.3.0) (2025-11-17)


### Features

* add new generator parser tool ([bc018e8](https://github.com/untrustedmodders/go-plugify/commit/bc018e8b626e3a56395bf7eaff264b540163219b))

## [1.2.4](https://github.com/untrustedmodders/go-plugify/compare/v1.2.3...v1.2.4) (2025-11-17)


### Bug Fixes

* add parser and move generators to separate folders ([84b1255](https://github.com/untrustedmodders/go-plugify/commit/84b1255462499cb21863f189449734a59f270d90))

## [1.2.3](https://github.com/untrustedmodders/go-plugify/compare/v1.2.2...v1.2.3) (2025-11-13)


### Bug Fixes

* add loaded flag and force gb on unload ([685247f](https://github.com/untrustedmodders/go-plugify/commit/685247f9c4c562c2ed98e56aafcbe012f9ebe5d3))
* expose PrintException ([35b5724](https://github.com/untrustedmodders/go-plugify/commit/35b57240243d5590242889e55fdb54f49c597354))

## [1.2.2](https://github.com/untrustedmodders/go-plugify/compare/v1.2.1...v1.2.2) (2025-09-13)


### Bug Fixes

* make plugin public ([def1a02](https://github.com/untrustedmodders/go-plugify/commit/def1a0280b5d1db9e07d756f6c8b03c76b4bd335))

## [1.2.1](https://github.com/untrustedmodders/go-plugify/compare/v1.2.0...v1.2.1) (2025-09-13)


### Bug Fixes

* remove missing methods ([76af23d](https://github.com/untrustedmodders/go-plugify/commit/76af23dc044f8a0fef8e53bc649ed08bb3f0c5c6))

## [1.2.0](https://github.com/untrustedmodders/go-plugify/compare/v1.1.10...v1.2.0) (2025-09-12)


### Features

* breaking changes ([b11d453](https://github.com/untrustedmodders/go-plugify/commit/b11d453870a1e9151592cdd32011c96777d46430))


### Bug Fixes

* update for a new plugify ([14b4221](https://github.com/untrustedmodders/go-plugify/commit/14b42212c218629e787e89575fef829be9173240))
* update for new plugify ([36caee2](https://github.com/untrustedmodders/go-plugify/commit/36caee2008929141370e049ef557a5e6e9dff9fd))

## [1.1.11](https://github.com/untrustedmodders/go-plugify/compare/v1.1.10...v1.1.11) (2025-09-06)


### Bug Fixes

* update for a new plugify ([14b4221](https://github.com/untrustedmodders/go-plugify/commit/14b42212c218629e787e89575fef829be9173240))

## [1.1.10](https://github.com/untrustedmodders/go-plugify/compare/v1.1.9...v1.1.10) (2025-08-29)


### Bug Fixes

* remove cgo nocallback ([e7100d8](https://github.com/untrustedmodders/go-plugify/commit/e7100d83fa713a6099ac9760850e2e1c1b95d097))

## [1.1.9](https://github.com/untrustedmodders/go-plugify/compare/v1.1.8...v1.1.9) (2025-08-17)


### Bug Fixes

* make plugin public ([bd0cb50](https://github.com/untrustedmodders/go-plugify/commit/bd0cb503fbe86118c5986008a9295d5d116fd54d))
* update IsPluginLoaded and IsModuleLoaded ([86bba60](https://github.com/untrustedmodders/go-plugify/commit/86bba60d25adcd7d6422feb3eb192f1c5266a72f))
* update readme and add missing permission ([1f544e9](https://github.com/untrustedmodders/go-plugify/commit/1f544e93ce0fdeffc5b06332c7072b216fef1c46))

## [1.1.8](https://github.com/untrustedmodders/go-plugify/compare/v1.1.7...v1.1.8) (2025-05-31)


### Bug Fixes

* add x86 support ([fdcf54d](https://github.com/untrustedmodders/go-plugify/commit/fdcf54d7f3af5dc51e29ec688b485dfbeaa0c6ee))
* make param array 64 bit for JitCall ([13b5df2](https://github.com/untrustedmodders/go-plugify/commit/13b5df221a490343a66b2a7a3e000aa95fd7eaf5))

## [1.1.7](https://github.com/untrustedmodders/go-plugify/compare/v1.1.6...v1.1.7) (2025-05-30)


### Bug Fixes

* add combability with 32 bit ([3ae92b0](https://github.com/untrustedmodders/go-plugify/commit/3ae92b0ecd2170e892e9a1d6a0378bc3fd72954b))
* hidden return on win arm and x86 ([08cd10d](https://github.com/untrustedmodders/go-plugify/commit/08cd10df93c2cb1f4182c64834f4384195dccead))

## [1.1.6](https://github.com/untrustedmodders/go-plugify/compare/v1.1.5...v1.1.6) (2025-05-15)


### Bug Fixes

* no return function ([55d14a3](https://github.com/untrustedmodders/go-plugify/commit/55d14a373531c3840a35c4a6f159e728cc192d7d))
* update generaotr ([dbcdef8](https://github.com/untrustedmodders/go-plugify/commit/dbcdef88cb33738c71e6d7041d34922dfef61da0))

## [1.1.5](https://github.com/untrustedmodders/go-plugify/compare/v1.1.4...v1.1.5) (2025-05-15)


### Bug Fixes

* remove char16 macro ([539173b](https://github.com/untrustedmodders/go-plugify/commit/539173b086d86aa6d3b1db2b9e1bed4a8ffccb6f))

## [1.1.4](https://github.com/untrustedmodders/go-plugify/compare/v1.1.3...v1.1.4) (2025-05-15)


### Bug Fixes

* apple clang missed uchar.h ([e0552e2](https://github.com/untrustedmodders/go-plugify/commit/e0552e2e6da440c700d38b82329d9bc40419f482))

## [1.1.3](https://github.com/untrustedmodders/go-plugify/compare/v1.1.2...v1.1.3) (2025-03-22)


### Bug Fixes

* add mutexes to function marshalling ([e6474d1](https://github.com/untrustedmodders/go-plugify/commit/e6474d160be90bd3f6b54de7a7f4e05aa803780f))
* function comparison ([413d668](https://github.com/untrustedmodders/go-plugify/commit/413d668bab8cc76a7e6f7d6daeb0f44c503b874e))

## [1.1.2](https://github.com/untrustedmodders/go-plugify/compare/v1.1.1...v1.1.2) (2025-03-17)


### Bug Fixes

* update native flags for new funcs ([5854d61](https://github.com/untrustedmodders/go-plugify/commit/5854d6141a64628c0a8a11d1e19f083d73e57f10))

## [1.1.1](https://github.com/untrustedmodders/go-plugify/compare/v1.1.0...v1.1.1) (2025-03-17)


### Bug Fixes

* remove GetBaseDir setter ([45c925f](https://github.com/untrustedmodders/go-plugify/commit/45c925f36fad17505de7239fe074be4ac42afaab))

## [1.1.0](https://github.com/untrustedmodders/go-plugify/compare/v1.0.5...v1.1.0) (2025-03-17)


### Features

* add path getters ([c2a36ad](https://github.com/untrustedmodders/go-plugify/commit/c2a36ad23768a6b14c2d7e865a1155c4c40af2e0))

## [1.0.5](https://github.com/untrustedmodders/go-plugify/compare/v1.0.4...v1.0.5) (2025-03-09)


### Bug Fixes

* add plugin context ([21d8963](https://github.com/untrustedmodders/go-plugify/commit/21d8963ae1b1382d42c1185ad1ee838269475924))

## [1.0.4](https://github.com/untrustedmodders/go-plugify/compare/v1.0.3...v1.0.4) (2025-03-09)


### Bug Fixes

* add missing callback ([e603e61](https://github.com/untrustedmodders/go-plugify/commit/e603e6121957e223086937b5555b7a27362d6bc7))

## [1.0.3](https://github.com/untrustedmodders/go-plugify/compare/v1.0.2...v1.0.3) (2025-03-09)


### Bug Fixes

* add missing includes to Linux ([caca727](https://github.com/untrustedmodders/go-plugify/commit/caca7278f5a7e84554c75382c660af01a41c89ca))

## [1.0.2](https://github.com/untrustedmodders/go-plugify/compare/v1.0.1...v1.0.2) (2025-03-09)


### Bug Fixes

* update go ([40c704e](https://github.com/untrustedmodders/go-plugify/commit/40c704e295733f1b94ffd324e52890f4a88a0ec7))

## [1.0.1](https://github.com/untrustedmodders/go-plugify/compare/v1.0.0...v1.0.1) (2025-03-09)


### Bug Fixes

* add sem versioning ([dafce37](https://github.com/untrustedmodders/go-plugify/commit/dafce37c4800140d082011beb027cb1799391f43))
