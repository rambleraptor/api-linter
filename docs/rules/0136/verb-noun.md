---
rule:
  aep: 136
  name: [core, '0136', verb-noun]
  summary: Custom methods should be named with the verb, then the noun.
permalink: /136/verb-noun
redirect_from:
  - /0136/verb-noun
---

# Custom methods: Verb and noun

This rule enforces that custom methods are named according to `VerbNoun`, as
mandated in [AEP-136][].

## Details

This is difficult to enforce without a dictionary (likely not worth it), so
this rule just attempts to catch common, easy to spot errors. It complains if:

- The method name is one word.

## Examples

### Single word method

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc Checkout(CheckoutRequest) returns (CheckoutResponse) {  // No noun.
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/books/*}:checkout"
    body: "*"
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc CheckoutBook(CheckoutBookRequest) returns (CheckoutBookResponse) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/books/*}:checkout"
    body: "*"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0136::verb-noun=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc Checkout(CheckoutRequest) returns (CheckoutResponse) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/books/*}:checkout"
    body: "*"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-136]: https://aep.dev/136
[aep.dev/not-precedent]: https://aep.dev/not-precedent
