---
rule:
  aep: 162
  name: [core, '0162', commit-request-name-behavior]
  summary: |
    Commit requests should annotate the `name` field with `google.api.field_behavior`.
permalink: /162/commit-request-name-behavior
redirect_from:
  - /0162/commit-request-name-behavior
---

# Commit requests: Name field behavior

This rule enforces that all `Commit` requests have
`google.api.field_behavior` set to `REQUIRED` on their `string name` field, as
mandated in [AEP-162][].

## Details

This rule looks at any message matching `Commit*Request` and complains if the
`name` field does not have a `google.api.field_behavior` annotation with a
value of `REQUIRED`.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message CommitBookRequest {
  // The `google.api.field_behavior` annotation should also be included.
  string name = 1 [
    (google.api.resource_reference).type = "library.googleapis.com/Book"
  ];
}
```

**Correct** code for this rule:

```proto
// Correct.
message CommitBookRequest {
  string name = 1 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference).type = "library.googleapis.com/Book"
  ];
}
```

## Disabling

If you need to violate this rule, use a leading comment above the field.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
message CommitBookRequest {
  // (-- api-linter: core::0162::commit-request-name-behavior=disabled
  //     aep.dev/not-precedent: We need to do this because reasons. --)
  string name = 1 [
    (google.api.resource_reference).type = "library.googleapis.com/Book"
  ];
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-162]: https://aep.dev/162
[aep.dev/not-precedent]: https://aep.dev/not-precedent
