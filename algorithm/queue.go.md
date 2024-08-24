---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  _extendedVerifiedWith:
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
    RuntimeError: bundler is not specified: algorithm/queue.go\n"
  code: "package algorithm\n\nimport \"errors\"\n\nvar ErrQueueEmpty = errors.New(\"\
    Queue: queue is empty\")\n\ntype Queue[T any] struct {\n\tdata []T\n}\n\nfunc\
    \ NewQueue[T any]() *Queue[T] {\n\treturn &Queue[T]{data: []T{}}\n}\n\nfunc (q\
    \ *Queue[T]) Empty() bool {\n\treturn q.Size() == 0\n}\n\nfunc (q *Queue[T]) Size()\
    \ int {\n\treturn len(q.data)\n}\n\nfunc (q *Queue[T]) Front() T {\n\tif q.Empty()\
    \ {\n\t\tpanic(ErrQueueEmpty)\n\t}\n\treturn q.data[0]\n}\n\nfunc (q *Queue[T])\
    \ Push(x T) {\n\tq.data = append(q.data, x)\n}\n\nfunc (q *Queue[T]) PushRange(v\
    \ []T) {\n\tfor i := 0; i < len(v); i++ {\n\t\tq.data = append(q.data, v[i])\n\
    \t}\n}\n\nfunc (q *Queue[T]) Pop() T {\n\tif q.Empty() {\n\t\tpanic(ErrQueueEmpty)\n\
    \t}\n\tres := q.data[0]\n\tq.data = q.data[1:]\n\treturn res\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/deque.go
  - algorithm/stack.go
  isVerificationFile: false
  path: algorithm/queue.go
  requiredBy:
  - main.go
  - algorithm/deque.go
  - algorithm/stack.go
  timestamp: '2024-08-24 14:32:30+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
documentation_of: algorithm/queue.go
layout: document
redirect_from:
- /library/algorithm/queue.go
- /library/algorithm/queue.go.html
title: algorithm/queue.go
---
