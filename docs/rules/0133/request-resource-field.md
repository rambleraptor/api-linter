---
rule:
  aep: 133
  name: [core, '0133', request-resource-field]
  summary: Create RPCs must have a field for the resource in the request.
permalink: /133/request-resource-field
redirect_from:
  - /0133/request-resource-field
---

# Create methods: Resource field

This rule enforces that all `Create` standard methods have a field in the
request message for the resource itself, as mandated in [AEP-133][].

## Details

This rule looks at any message matching `Create*Request` and complains if there
is no field of the resource's type with the expected field name.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
// `Book book` is missing.
message CreateBookRequest {
  string publisher = 1;
  string book_id = 3;
}
```

```proto
// Incorrect.
message CreateBookRequest {
  bytes parent = 1;
  Book payload = 2;  // Field name should be `book`.
  string book_id = 3;
}
```

**Correct** code for this rule:

```proto
// Correct.
message CreateBookRequest {
  string parent = 1;
  Book book = 2;
  string book_id = 3;
}
```

## Disabling

If you need to violate this rule, use a leading comment above the message (if
the resource field is missing) or above the field (if it is improperly named).
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0133::request-resource-field=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
message CreateBookRequest {
  string publisher = 1;
  Book payload = 2;
  string book_id = 3;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-133]: https://aep.dev/133
[aep.dev/not-precedent]: https://aep.dev/not-precedent
