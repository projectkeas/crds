# Project Keas - Custom Resource Definitions

This repository contains the Custom Resource Definitions that are used within the Keas project. The types exposed are:

- keas.io/v1alpha1
  - EventType
  - IngestionPolicy

## keas.io/v1alpha1

### EventType

The `EventType` resource allows you to define a versionable type that is accepted by the ingestion engine. This helps to ensure the consistency of data within the system.

|Field|Description|Required|
|---|---|---|
|owner|The owner of the type, eg: ProjectKeas. Must be match regex: `^[A-z\-]{3,63}$`|true|
|version|The version of the event|true
|schema|A json schema that will be used to validate the incoming event|true|
|type|The name of the event, eg: Repository. Must be match regex: `^[A-z\-]{3,63}$`. This will be case-sensitive.|true
|sources|The allowed list of sources for this event type. This will be case-sensitive.|false|
|subTypes|The allowed subtypes for the event, eg: Created/Updated/Deleted. This will be case-sensitive.|false|
|description|A friendly description of the type|false|

Example:

```yaml
apiVersion: keas.io/v1alpha1
kind: EventType
metadata:
  name: TestType
spec:
  description: "This is a type for testing"
  version: '0.0.1'
  owner: Test
  type: TestType
  sources:
    - Gitlab
    - Github
  subTypes:
    - Created
    - Updated
    - Deleted
  schema: |
    {
        "$id": "https://example.com/person.schema.json",
        "$schema": "https://json-schema.org/draft/2020-12/schema",
        "title": "Person",
        "type": "object",
        "properties": {
            "firstName": {
            "type": "string",
            "description": "The person's first name."
            },
            "lastName": {
            "type": "string",
            "description": "The person's last name."
            },
            "age": {
            "description": "Age in years which must be equal to or greater than zero.",
            "type": "integer",
            "minimum": 0
            }
        }
    }
```

### IngestionPolicy

The `IngestionPolicy` type governs whether a given event is accepted by the system for storage and how long the event should be stored for. There are two parts to the policy that can be used to control the ingestion:

- `allow`: Determines whether or not the system will store the event. Default: true
- `ttl`: Determines the length of time (in seconds) that the event be stored for before deletion. Default: -1 (forever)

The input to the policy is the complete request from the ingestion web request (ie: contains both the metadata and payload). All ingestion policies are evaluated for all requests.

|Field|Description|Required|
|---|---|---|
|policy|The rego representation of the policy|true|
|defaults|Controls the defaults for the policy. Default: `allow: true` and `ttl: -1`|false|
|description|A description of what the policy does|false|

Example:

```yaml
apiVersion: keas.io/v1alpha1
kind: IngestionPolicy
metadata:
  name: test2
spec:
  description: Simple policy to only ingest
  defaults:
    allow: true
    ttl: -1
  policy: |
    # When: defaults.allow = true
    # Do not store the event when the last name is "world"
    allow = false {
        input.payload.lastName == "world"
    }
    # When: defaults.allow = false
    # All the event to be stored when the last name is not "world"
    allow { 
        input.payload.lastName != "world"
    }
    # Sets the ttl to 60 seconds when the first name is "hello"
    ttl = 60 {
        input.payload.firstName == "hello"
    }
```

## Updating the CRDs

In order to update the CRDs run the `generate.sh` script. This will remove all of the generated types and re-generate them according to types defined in each package.
