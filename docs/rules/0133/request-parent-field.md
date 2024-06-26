---
rule:
  aep: 133
  name: [core, '0133', request-parent-field]
  summary: Create RPCs must have a `parent` field in the request.
permalink: /133/request-parent-field
redirect_from:
  - /0133/request-parent-field
---

# Create methods: Parent field

This rule enforces that all `Create` standard methods have a `string parent`
field in the request message, as mandated in [AEP-133][].

## Details

This rule looks at any message matching `Create*Request` and complains if
the `parent` field has any type other than `string`.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message GetBookRequest {
  // Field type should be `string`.
  bytes parent = 1;
  Book book = 2;
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

If you need to violate this rule, use a leading comment above the field.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
message CreateBookRequest {
  // (-- api-linter: core::0133::request-parent-field=disabled
  //     aep.dev/not-precedent: We need to do this because reasons. --)
  bytes parent = 1;
  Book book = 2;
  string book_id = 3;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-133]: https://aep.dev/133
[aep.dev/not-precedent]: https://aep.dev/not-precedent
