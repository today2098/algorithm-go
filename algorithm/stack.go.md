---
data:
  _extendedDependsOn:
  - icon: ':warning:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':warning:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':warning:'
    path: main.go
    title: main.go
  _extendedRequiredBy:
  - icon: ':warning:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':warning:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':warning:'
    path: main.go
    title: main.go
  _extendedVerifiedWith: []
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':warning:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: algorithm/stack.go\n"
  code: "package algorithm\n\nimport \"errors\"\n\nvar ErrStackEmpty = errors.New(\"\
    Stack: stack is empty\")\n\ntype Stack[T any] struct {\n\tdata []T\n}\n\nfunc\
    \ NewStack[T any]() *Stack[T] {\n\treturn &Stack[T]{data: []T{}}\n}\n\nfunc (s\
    \ *Stack[T]) Empty() bool {\n\treturn s.Size() == 0\n}\n\nfunc (s *Stack[T]) Size()\
    \ int {\n\treturn len(s.data)\n}\n\nfunc (s *Stack[T]) Top() T {\n\tif s.Empty()\
    \ {\n\t\tpanic(ErrStackEmpty)\n\t}\n\treturn s.data[len(s.data)-1]\n}\n\nfunc\
    \ (s *Stack[T]) Push(x T) {\n\ts.data = append(s.data, x)\n}\n\nfunc (s *Stack[T])\
    \ PushRange(v []T) {\n\tfor i := 0; i < len(v); i++ {\n\t\ts.data = append(s.data,\
    \ v[i])\n\t}\n}\n\nfunc (s *Stack[T]) Pop() T {\n\tif s.Empty() {\n\t\tpanic(ErrStackEmpty)\n\
    \t}\n\tres := s.data[len(s.data)-1]\n\ts.data = s.data[:len(s.data)-1]\n\treturn\
    \ res\n}\n"
  dependsOn:
  - main.go
  - algorithm/deque.go
  - algorithm/queue.go
  isVerificationFile: false
  path: algorithm/stack.go
  requiredBy:
  - main.go
  - algorithm/deque.go
  - algorithm/queue.go
  timestamp: '2024-08-24 13:59:30+09:00'
  verificationStatus: LIBRARY_NO_TESTS
  verifiedWith: []
documentation_of: algorithm/stack.go
layout: document
redirect_from:
- /library/algorithm/stack.go
- /library/algorithm/stack.go.html
title: algorithm/stack.go
---
