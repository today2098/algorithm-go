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
    path: algorithm/stack/aoj-ALDS1_3_A.test.go
    title: algorithm/stack/aoj-ALDS1_3_A.test.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack/stack.go
    title: algorithm/stack/stack.go
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
    path: algorithm/stack/stack.go
    title: algorithm/stack/stack.go
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
    RuntimeError: bundler is not specified: algorithm/queue/queue.go\n"
  code: "package queue\n\nimport \"errors\"\n\nvar ErrQueueEmpty = errors.New(\"Queue:\
    \ queue is empty\")\n\ntype Queue[T any] struct {\n\tData []T\n}\n\nfunc NewQueue[T\
    \ any]() *Queue[T] {\n\treturn &Queue[T]{Data: []T{}}\n}\n\nfunc (q *Queue[T])\
    \ Empty() bool {\n\treturn q.Size() == 0\n}\n\nfunc (q *Queue[T]) Size() int {\n\
    \treturn len(q.Data)\n}\n\nfunc (q *Queue[T]) Front() T {\n\tif q.Empty() {\n\t\
    \tpanic(ErrQueueEmpty)\n\t}\n\treturn q.Data[0]\n}\n\nfunc (q *Queue[T]) Push(x\
    \ T) {\n\tq.Data = append(q.Data, x)\n}\n\nfunc (q *Queue[T]) PushRange(v []T)\
    \ {\n\tfor i := 0; i < len(v); i++ {\n\t\tq.Data = append(q.Data, v[i])\n\t}\n\
    }\n\nfunc (q *Queue[T]) Pop() T {\n\tif q.Empty() {\n\t\tpanic(ErrQueueEmpty)\n\
    \t}\n\tres := q.Data[0]\n\tq.Data = q.Data[1:]\n\treturn res\n}\n"
  dependsOn:
  - main.go
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/stack/stack.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/deque/deque.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
  isVerificationFile: false
  path: algorithm/queue/queue.go
  requiredBy:
  - main.go
  - algorithm/stack/stack.go
  - algorithm/deque/deque.go
  timestamp: '2024-08-24 17:41:38+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-ITP1_1_A.test.go
  - algorithm/stack/aoj-ALDS1_3_A.test.go
  - algorithm/queue/aoj-ALDS1_3_B.test.go
  - algorithm/deque/aoj-ALDS1_3_C.test.go
documentation_of: algorithm/queue/queue.go
layout: document
redirect_from:
- /library/algorithm/queue/queue.go
- /library/algorithm/queue/queue.go.html
title: algorithm/queue/queue.go
---
