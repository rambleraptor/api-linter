---
rule:
  aep: 165
  name: [core, '0165', request-parent-behavior]
  summary: |
    Purge requests should annotate the `parent` field with `google.api.field_behavior`.
permalink: /165/request-parent-behavior
redirect_from:
  - /0165/request-parent-behavior
---

# Purge requests: Parent field behavior

This rule enforces that all `Purge` requests have
`google.api.field_behavior` set to `REQUIRED` on their `string parent` field, as
mandated in [AEP-165][].

## Details

This rule looks at any message matching `Purge*Request` and complains if the
`parent` field does not have a `google.api.field_behavior` annotation with a
value of `REQUIRED`.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message PurgeBooksRequest {
  // The `google.api.field_behavior` annotation should also be included.
  string parent = 1 [
    (google.api.resource_reference).child_type = "library.googleapis.com/Book"
  ];

  string filter = 2 [(google.api.field_behavior) = REQUIRED];

  bool force = 3;
}
```

**Correct** code for this rule:

```proto
// Correct.
message PurgeBooksRequest {
  string parent = 1 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference).child_type = "library.googleapis.com/Book"
  ];

  string filter = 2 [(google.api.field_behavior) = REQUIRED];

  bool force = 3;
}
```

## Disabling

If you need to violate this rule, use a leading comment above the field.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
message PurgeBooksRequest {
  // (-- api-linter: core::0165::request-parent-behavior=disabled
  //     aep.dev/not-precedent: We need to do this because reasons. --)
  string parent = 1 [
    (google.api.resource_reference).child_type = "library.googleapis.com/Book"
  ];

  string filter = 2 [(google.api.field_behavior) = REQUIRED];

  bool force = 3;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-165]: https://aep.dev/165
[aep.dev/not-precedent]: https://aep.dev/not-precedent
