---
rule:
  aep: 203
  name: [core, '0203', field-behavior-required]
  summary: |
    Field behavior is required, and must have one of OUTPUT_ONLY, REQUIRED, or
    OPTIONAL.
permalink: /203/field-behavior-required
redirect_from:
  - /0203/field-behavior-required
---

# Field Behavior Required

This rule enforces that each field in a message used in a request has a
`google.api.field_behavior` annotation with valid values, as mandated by
[AEP-203][].

## Details

This rule looks at all fields except those nested inside a OneOf and ensures
they have a `google.api.field_behavior` annotation. It also verifies that they
have at least one of the options `OUTPUT_ONLY`, `REQUIRED`, or `OPTIONAL`: all
fields must fall into one of these categories.

Although all request messages **must** be annotated, this linter only verifies
messages that are in the same package as some upstream protos (e.g. common
protos) may be difficult to modify.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message Book {
  string name = 1;

  // No field behavior
  optional string title = 2;
}
```

**Correct** code for this rule:

```proto
// Correct.
message Book {
  string name = 1;

  string title = 2 [(google.api.field_behavior) = REQUIRED];
}
```

## Disabling

If you need to violate this rule, use a leading comment above the field.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
message Book {
  string name = 1;

  // Required. The title of the book.
  // (-- api-linter: core::0203::field-behavior-required=disabled
  //     aep.dev/not-precedent: We need to do this because reasons. --)
  optional string title = 2;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-203]: https://aep.dev/203
[aep.dev/not-precedent]: https://aep.dev/not-precedent
