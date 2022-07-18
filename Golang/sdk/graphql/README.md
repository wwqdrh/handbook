核心就是

```go
schema := schemabuilder.NewSchema

schema.Query.FieldFunc("[name]", func...)

schema.Mutation.FieldFunc("[name]", func...)

schema.Object.FieldFunc("[name]", func...)
```