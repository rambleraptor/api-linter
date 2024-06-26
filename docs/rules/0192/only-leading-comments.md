---
rule:
  aep: 192
  name: [core, '0192', only-leading-comments]
  summary: All public comments should be leading comments.
permalink: /192/only-leading-comments
redirect_from:
  - /0192/only-leading-comments
---

# Only leading comments

This rule enforces that every descriptor in every proto file has a public
comment information only in leading comment (not trailing comments or detached
comments), as mandated in [AEP-192][].

## Details

This rule looks at each descriptor in each proto file (exempting oneofs and the
file itself) and complains if _public_ comments are either trailing or
detached. Internal comments are not considered.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
// BEGIN LIBRARY SECTION  <-- detached comment; shows up in docs.

// A representation of a book.
message Book {
  // The resource name of the book.
  string name = 1;
}
```

**Correct** code for this rule:

```proto
// Correct.
// (-- BEGIN LIBRARY SECTION --)

// A representation of a book.
message Book {
  // The resource name of the book.
  string name = 1;
}
```

## Disabling

If you need to violate this rule, use a leading comment above the descriptor
(and revel in the irony). Remember to also include an [aep.dev/not-precedent][]
comment explaining why.

```proto
// BEGIN LIBRARY SECTION  <-- detached comment; shows up in docs.

// A representation of a book.
// (-- api-linter: core::0192::only-leading-comments=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
message Book {
  // The resource name of the book.
  string name = 1;
}
```

[aep-192]: https://aep.dev/192
[aep.dev/not-precedent]: https://aep.dev/not-precedent
