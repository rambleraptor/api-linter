---
rule:
  aep: 146
  name: [core, '0146', any]
  summary: Avoid `google.protobuf.Any` fields.
permalink: /146/any
redirect_from:
  - /0146/any
---

# Any

This rule discourages the use of `google.protobuf.Any`, as described in
[AEP-146][].

## Details

This rule complains if it sees a `google.protobuf.Any` field. Common packages
(such as `google.api` or `google.longrunning`) are excluded.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message Book {
  // google.protobuf.Any is discouraged.
  google.protobuf.Any contents = 1;
}
```

**Correct** code for this rule:

The correct code is likely to vary substantially by use case. See [AEP-146][]
for details and tradeoffs of various approaches for generic fields.

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0146::any=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
message Book {
  google.protobuf.Any contents = 1;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-146]: https://aep.dev/146
[aep.dev/not-precedent]: https://aep.dev/not-precedent
