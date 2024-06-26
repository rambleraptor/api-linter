---
rule:
  aep: 151
  name: [core, '0151', lro-response-reachable]
  summary: LRO response messages must be reachable.
permalink: /151/lro-response-reachable
redirect_from:
  - /0151/lro-response-reachable
---

# LRO: Reachable response types

This rule enforces that methods returning long-running operations define their
response messages in the same file or a directly-imported file, as mandated in
[AEP-151][].

## Details

This rule looks at any method with a return type of
`google.longrunning.Operation`, and complains if the message designated by the
`response_type` field are not defined in the same file, or a file directly
imported by the file.

Because these message names are strings, and a string reference does not
require an `import` statement, defining the response types elsewhere can cause
problems for tooling. To prevent this, and also to maintain consistency with
the file layout in [AEP-191][], the linter complains if the message is not
defined in the same file or a file directly imported by the file.

## Examples

**Incorrect** code for this rule:

In `library_service.proto`:

```proto
// Incorrect.
service Library {
  rpc WriteBook(WriteBookRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=publishers/*/books}:write"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "WriteBookResponse"
      metadata_type: "WriteBookMetadata"
    };
  }
}
```

In `operations.proto`:

```proto
// Incorrect.
message WriteBookResponse {
  // Should be in the same file or directly imported.
}

message WriteBookMetadata {
  // Should be in the same file or directly imported.
}
```

**Correct** code for this rule:

Same file:

```proto
// Correct.
service Library {
  rpc WriteBook(WriteBookRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=publishers/*/books}:write"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "WriteBookResponse"
      metadata_type: "WriteBookMetadata"
    };
  }
}

// Later in the file...
message WriteBookResponse {
  // ...
}

message WriteBookMetadata {
  // ...
}
```

Separate files:

In `library_service.proto`:

```proto
// Correct.
import "operations.proto";

service Library {
  rpc WriteBook(WriteBookRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v1/{name=publishers/*/books}:write"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "WriteBookResponse"
      metadata_type: "WriteBookMetadata"
    };
  }
}
```

In `operations.proto`:

```proto
// Correct.
message WriteBookResponse {
  // ...
}

message WriteBookMetadata {
  // ...
}
```

## Disabling

If you need to violate this rule, use a leading comment above the method.
Remember to also include an [aep.dev/not-precedent][] comment explaining why.

```proto
// (-- api-linter: core::0151::lro-response-reachable=disabled
//     aep.dev/not-precedent: We need to do this because reasons. --)
rpc WriteBook(WriteBookRequest) returns (google.longrunning.Operation) {
  option (google.api.http) = {
    post: "/v1/{name=publishers/*/books}:write"
    body: "*"
  };
  option (google.longrunning.operation_info) = {
    response_type: "google.protobuf.Empty"
    metadata_type: "WriteBookMetadata"
  };
}
```

If you need to violate this rule for an entire file, place the comment at the
top of the file.

[aep-151]: https://aep.dev/151
[aep-191]: https://aep.dev/191
[aep.dev/not-precedent]: https://aep.dev/not-precedent
