---
data:
  _extendedDependsOn: []
  _extendedRequiredBy: []
  _extendedVerifiedWith: []
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':warning:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: main.go\n"
  code: "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, world!\"\
    )\n}\n"
  dependsOn: []
  isVerificationFile: false
  path: main.go
  requiredBy: []
  timestamp: '2024-08-24 13:28:00+09:00'
  verificationStatus: LIBRARY_NO_TESTS
  verifiedWith: []
documentation_of: main.go
layout: document
redirect_from:
- /library/main.go
- /library/main.go.html
title: main.go
---
