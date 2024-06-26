---
rule:
  aep: 148
  name: [core, '0148', human-names]
  summary: Avoid imprecise terms for human names.
permalink: /148/human-names
redirect_from:
  - /0148/human-names
---

# Human names

This rule encourages terms for human names (`given_name` and `family_name`)
that are more accurate across cultures, as mandated in [AEP-148][].

## Details

This rule looks for fields named `first_name` and `last_name`, and complains if
it finds them, suggesting the use of `given_name` and `family_name`
(respectively) instead.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
message Human {
  string first_name = 1;  // Should be `given_name`.
  string last_name = 2;   // Should be `family_name`
}
```

**Correct** code for this rule:

```proto
// Correct.
message Human {
  string given_name = 1;
  string family_name = 2;
}
```

## Disabling

If you need to violate this rule, use a leading comment above the field or its
enclosing message. Remember to also include an [aep.dev/not-precedent][]
comment explaining why.

```proto
// (-- api-linter: core::0148::human-names=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
message Human {
  string first_name = 1;
  string last_name = 2;
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-148]: https://aep.dev/148
[aep.dev/not-precedent]: https://aep.dev/not-precedent
