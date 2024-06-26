---
rule:
  aep: 231
  name: [core, '0231', request-parent-field]
  summary: Batch Get RPCs must have a `parent` field in the request.
permalink: /231/request-parent-field
redirect_from:
  - /0231/request-parent-field
---

# Batch Get methods: Parent field

This rule enforces that all `BatchGet` methods have a `string parent` field in
the request message, as mandated in [AEP-231][].

## Details

This rule looks at any message matching `BatchGet*Request` and complains if
either the `parent` field is missing, or if it has any type other than
`string`.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message BatchGetBooksRequest {
  string publisher = 1;  // Field name should be `parent`.

  repeated string names = 2 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference).type = "library.googleapis.com/Book"
  ];
}
```

```proto
// Incorrect.
message BatchGetBooksRequest {
  bytes parent = 1;  // Field type should be `string`.

  repeated string names = 2 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference).type = "library.googleapis.com/Book"
  ];
}
```

**Correct** code for this rule:

```proto
// Correct.
message BatchGetBooksRequest {
  string parent = 1 [
    (google.api.resource_reference).child_type = "library.googleapis.com/Book"
  ];

  repeated string names = 2 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference).type = "library.googleapis.com/Book"
  ];
}
```

## Disabling

If you need to violate this rule, use a leading comment above the message (if
the `parent` field is missing) or above the field (if it is the wrong type).
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0231::request-parent-field=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
message BatchGetBooksRequest {
  string publisher = 1;  // Field name should be `parent`.

  repeated string names = 2 [
    (google.api.field_behavior) = REQUIRED,
    (google.api.resource_reference).type = "library.googleapis.com/Book"
  ];
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-231]: https://aep.dev/231
[aep.dev/not-precedent]: https://aep.dev/not-precedent
