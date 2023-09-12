使用`github.com/samsarahq/thunder`包构建graphql service

核心就是

```go
schema := schemabuilder.NewSchema

schema.Query.FieldFunc("[name]", func...)

schema.Mutation.FieldFunc("[name]", func...)

schema.Object.FieldFunc("[name]", func...)
```