---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/aoj-ALDS1_3_C.test.go
    title: algorithm/deque/aoj-ALDS1_3_C.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/deque.go
    title: algorithm/deque/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/aoj-ALDS1_3_B.test.go
    title: algorithm/queue/aoj-ALDS1_3_B.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/queue.go
    title: algorithm/queue/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/aoj-ALDS1_3_A.test.go
    title: algorithm/stack/aoj-ALDS1_3_A.test.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/deque.go
    title: algorithm/deque/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/queue.go
    title: algorithm/queue/queue.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque/aoj-ALDS1_3_C.test.go
    title: algorithm/deque/aoj-ALDS1_3_C.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue/aoj-ALDS1_3_B.test.go
    title: algorithm/queue/aoj-ALDS1_3_B.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/aoj-ALDS1_3_A.test.go
    title: algorithm/stack/aoj-ALDS1_3_A.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: algorithm/stack/stack.go\n"
  code: "package stack\n\nimport \"errors\"\n\nvar ErrStackEmpty = errors.New(\"Stack:\
    \ stack is empty\")\n\ntype Stack[T any] struct {\n\tData []T\n}\n\nfunc NewStack[T\
    \ any]() *Stack[T] {\n\treturn &Stack[T]{Data: []T{}}\n}\n\nfunc (s *Stack[T])\
    \ Empty() bool {\n\treturn s.Size() == 0\n}\n\nfunc (s *Stack[T]) Size() int {\n\
    \treturn len(s.Data)\n}\n\nfunc (s *Stack[T]) Top() T {\n\tif s.Empty() {\n\t\t\
    panic(ErrStackEmpty)\n\t}\n\treturn s.Data[len(s.Data)-1]\n}\n\nfunc (s *Stack[T])\
    \ Push(x T) {\n\ts.Data = append(s.Data, x)\n}\n\nfunc (s *Stack[T]) PushRange(v\
    \ []T) {\n\tfor i := 0; i < len(v); i++ {\n\t\ts.Data = append(s.Data, v[i])\n\
    \t}\n}\n\nfunc (s *Stack[T]) Pop() T {\n\tif s.Empty() {\n\t\tpanic(ErrStackEmpty)\n\
    \t}\n\tres := s.Data[len(s.Data)-1]\n\ts.Data = s.Data[:len(s.Data)-1]\n\treturn\
    \ res\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/queue/queue.go
  - algorithm/deque/deque.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
  isVerificationFile: false
  path: algorithm/stack/stack.go
  requiredBy:
  - main.go
  - algorithm/queue/queue.go
  - algorithm/deque/deque.go
  timestamp: '2024-08-24 17:41:38+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
documentation_of: algorithm/stack/stack.go
layout: document
redirect_from:
- /library/algorithm/stack/stack.go
- /library/algorithm/stack/stack.go.html
title: algorithm/stack/stack.go
---
