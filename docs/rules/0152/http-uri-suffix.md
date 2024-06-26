---
rule:
  aep: 152
  name: [core, '0152', http-uri-suffix]
  summary: Run methods must have the correct URI suffix
permalink: /152/http-uri-suffix
redirect_from:
  - /0152/http-uri-suffix
---

# Run methods: URI suffix

This rule enforces that `Run` methods include the `:run` suffix
in the REST URI, as mandated in [AEP-152][].

## Details

This rule looks at any method whose name starts with `Run`, and
complains if the HTTP URI does not end with `:run`.

## Examples

**Incorrect** code for this rule:

```proto
// Incorrect.
rpc RunWriteBookJob(RunWriteBookJobRequest) returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/writeBookJobs/*}:execute" // Should end with `:run`.
    body: "*"
  };
  option (google.longrunning.operation_info) = {
    response_type: "RunWriteBookJobResponse"
    metadata_type: "RunWriteBookJobMetadata"
  };
}
```

**Correct** code for this rule:

```proto
// Correct.
rpc RunWriteBookJob(RunWriteBookJobRequest) returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/writeBookJobs/*}:run"
    body: "*"
  };
  option (google.longrunning.operation_info) = {
    response_type: "RunWriteBookJobResponse"
    metadata_type: "RunWriteBookJobMetadata"
  };
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0152::http-uri-suffix=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc RunWriteBookJob(RunWriteBookJobRequest) returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/writeBookJobs/*}:execute"
    body: "*"
  };
  option (google.longrunning.operation_info) = {
    response_type: "RunWriteBookJobResponse"
    metadata_type: "RunWriteBookJobMetadata"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-152]: https://aep.dev/152
[aep.dev/not-precedent]: https://aep.dev/not-precedent
