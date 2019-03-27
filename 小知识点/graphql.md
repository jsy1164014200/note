# graphql

[github v4 用的graphql ](https://developer.github.com/v4/)

GraphQL 也计划引入除了 `query`, `mutation` 以外的第三种操作符 `subscription`，以便于直接接受服务器推送数据。

## 查询和更新

### 1. 字段 Fields

```js
{
  hero {
    name
    # 查询可以有备注！
    friends {
      name
    }
  }
}

{
  "data": {
    "hero": {
      "name": "R2-D2",
      "friends": [
        {
          "name": "Luke Skywalker"
        },
        {
          "name": "Han Solo"
        },
        {
          "name": "Leia Organa"
        }
      ]
    }
  }
}
```

### 2. 参数(Arguments)

每一个字段和嵌套对象都能有自己的一组参数，从而使得 GraphQL 可以完美替代多次 API 获取请求。甚至你也可以给 标量（scalar）字段传递参数，用于实现服务端的一次转换，而不用每个客户端分别转换。

```js
{
    human(id: "1000") {
    	name
    	height(unit: FOOT)
  }
}

{
  "data": {
    "human": {
      "name": "Luke Skywalker",
      "height": 5.6430448
    }
  }
}
```

### 3. 别名 aliases

```js
{
  empireHero: hero(episode: EMPIRE) {
    name
  }
  jediHero: hero(episode: JEDI) {
    name
  }
}

{
  "data": {
    "empireHero": {
      "name": "Luke Skywalker"
    },
    "jediHero": {
      "name": "R2-D2"
    }
  }
}
```

上例中，两个 `hero` 字段将会存在冲突，但是因为我们可以将其另取一个别名，我们也就可以在一次请求中得到两个结果。

### 4. 片段（类似函数调用）

片段使你能够组织一组字段，然后在需要它们的的地方引入。

```js
{
  leftComparison: hero(episode: EMPIRE) {
    ...comparisonFields
  }
  rightComparison: hero(episode: JEDI) {
    ...comparisonFields
  }
}

fragment comparisonFields on Character {
  name
  appearsIn
  friends {
    name
  }
}
```

##### 可以在函数中使用 变量

```js
# { "graphiql": true, "variables": { "episode": JEDI } }
query HeroNameAndFriends($episode: Episode) {
  hero(episode: $episode) {
    name
    friends {
      name
    }
  }
}
```

变量定义看上去像是上述查询中的 `($episode: Episode)`。其工作方式跟类型语言中函数的参数定义一样。它以列出所有变量，变量前缀必须为 `$`，后跟其类型，本例中为 `Episode`。

变量定义可以是可选的或者必要的。上例中，`Episode` 后并没有 `!`，因此其是可选的。但是如果你传递变量的字段要求非空参数，那变量一定是必要的。

变量也可以带默认值

```js
# $first 就是一个变量
query HeroComparison($first: Int = 3) {
  leftComparison: hero(episode: EMPIRE) {
    ...comparisonFields
  }
  rightComparison: hero(episode: JEDI) {
    ...comparisonFields
  }
}

fragment comparisonFields on Character {
  name
  friendsConnection(first: $first) {
    totalCount
    edges {
      node {
        name
      }
    }
  }
}
```





### 5. 操作名称 operation name

这之前，我们都使用了简写句法，省略了 `query` 关键字和查询名称，但是生产中使用这些可以使我们代码减少歧义。

下面的示例包含了作为**操作类型**的关键字 `query` 以及**操作名称** `HeroNameAndFriends`

```js
query HeroNameAndFriends {
  hero {
    name
    friends {
      name
    }
  }
}

```

**操作类型**可以是 *query*、*mutation* 或 *subscription*，描述你打算做什么类型的操作。操作类型是必需的，除非你使用查询简写语法，在这种情况下，你无法为操作提供名称或变量定义。

### 6. 使用指令

- `@include(if: Boolean)` 仅在参数为 `true` 时，包含此字段。
- `@skip(if: Boolean)` 如果参数为 `true`，跳过此字段。

```js
query Hero($episode: Episode, $withFriends: Boolean!) {
  hero(episode: $episode) {
    name
    friends @include(if: $withFriends) {
      name
    }
  }
}

# 如果 withFriends 是 true
{
  "data": {
    "hero": {
      "name": "R2-D2",
      "friends": [
        {
          "name": "Luke Skywalker"
        },
        {
          "name": "Han Solo"
        },
        {
          "name": "Leia Organa"
        }
      ]
    }
  }
}

```

### 7. 更新(Mutations)

```js
mutation CreateReviewForEpisode($ep: Episode!, $review: ReviewInput!) {
  createReview(episode: $ep, review: $review) {
    stars
    commentary
  }
}

# 参数
{
  "ep": "JEDI",
  "review": {
    "stars": 5,
    "commentary": "This is a great movie!"
  }
}

# 结果
{
  "data": {
    "createReview": {
      "stars": 5,
      "commentary": "This is a great movie!"
    }
  }
}
```



### 8. 内联片段 inline fragments

```js
query HeroForEpisode($ep: Episode!) {
  hero(episode: $ep) {
    name
    ... on Droid {
      primaryFunction
    }
    ... on Human {
      height
    }
  }
}

因为第一个片段标注为 ... on Droid，primaryFunction 仅在 hero 返回的 Character 为 Droid 类型时才会执行。同理适用于 Human 类型的 height 字段。
{
  "data": {
    "hero": {
      "name": "R2-D2",
      "primaryFunction": "Astromech"
    }
  }
}
```



### 9. 元字段 Meta fields

```js
{
  search(text: "an") {
    __typename
    ... on Human {
      name
    }
    ... on Droid {
      name
    }
    ... on Starship {
      name
    }
  }
}

# 元字段 __typename
{
  "data": {
    "search": [
      {
        "__typename": "Human",
        "name": "Han Solo"
      },
      {
        "__typename": "Human",
        "name": "Leia Organa"
      },
      {
        "__typename": "Starship",
        "name": "TIE Advanced x1"
      }
    ]
  }
}
```



## Schema 和 类型

> GraphQL 类型系统

> schema 模式，可以理解成一个集合中的 文档，一个 表中的一行

### 1. 对象类型和字段

```js
type Character {
  name: String!
  appearsIn: [Episode!]!
}
```

- `Character` 是一个 **GraphQL 对象类型**，表示其是一个拥有一些字段的类型。你的 schema 中的大多数类型都会是对象类型。
- `name` 和 `appearsIn` 是 `Character` 类型上的**字段**。这意味着在一个操作 `Character`类型的 GraphQL 查询中的任何部分，都只能出现 `name` 和 `appearsIn` 字段。
- `String` 是内置的**标量**类型之一 —— 标量类型是解析到单个标量对象的类型，无法在查询中对它进行次级选择。后面我们将细述标量类型。
- `String!` 表示这个字段是**非空的**，GraphQL 服务保证当你查询这个字段后总会给你返回一个值。在类型语言里面，我们用一个感叹号来表示这个特性。
- `[Episode!]!` 表示一个 `Episode` **数组**。因为它也是**非空的**，所以当你查询 `appearsIn` 字段的时候，你也总能得到一个数组（零个或者多个元素）。且由于 `Episode!` 也是**非空的**，你总是可以预期到数组中的每个项目都是一个 `Episode` 对象。

#### 字段上的参数

```js
type Starship {
  id: ID!
  name: String!
  length(unit: LengthUnit = METER): Float
}
```



### 2. 查询和变更类型

你的 schema 中大部分的类型都是普通对象类型，但是一个 schema 内有两个特殊类型：

```graphql
schema {
  query: Query
  mutation: Mutation
}
```

每一个 GraphQL 服务都有一个 `query` 类型，可能有一个 `mutation` 类型。这两个类型和常规对象类型无差，但是它们之所以特殊，是因为它们定义了每一个 GraphQL 查询的**入口**。因此如果你看到一个像这样的查询：

```
query {
  hero {
    name
  }
  droid(id: "2000") {
    name
  }
}
```

```
{
  "data": {
    "hero": {
      "name": "R2-D2"
    },
    "droid": {
      "name": "C-3PO"
    }
  }
}
```

那表示这个 GraphQL 服务需要一个 `Query` 类型，且其上有 `hero` 和 `droid` 字段：

```graphql
type Query {
  hero(episode: Episode): Character
  droid(id: ID!): Droid
}
```



### 3. 标量类型

- `Int`：有符号 32 位整数。
- `Float`：有符号双精度浮点值。
- `String`：UTF‐8 字符序列。
- `Boolean`：`true` 或者 `false`。
- `ID`：ID 标量类型表示一个唯一标识符，通常用以重新获取对象或者作为缓存中的键。ID 类型使用和 String 一样的方式序列化；然而将其定义为 ID 意味着并不需要人类可读型。

大部分的 GraphQL 服务实现中，都有自定义标量类型的方式。例如，我们可以定义一个 `Date` 类型：

```graphql
scalar Date
```

然后就取决于我们的实现中如何定义将其序列化、反序列化和验证。例如，你可以指定 `Date` 类型应该总是被序列化成整型时间戳，而客户端应该知道去要求任何 date 字段都是这个格式。

### 4. 枚举类型

下面是一个用 GraphQL schema 语言表示的 enum 定义：

```graphql
enum Episode {
  NEWHOPE
  EMPIRE
  JEDI
}
```

### 5. 列表和非空

对象类型、标量以及枚举是 GraphQL 中你唯一可以定义的类型种类。但是当你在 schema 的其他部分使用这些类型时，或者在你的查询变量声明处使用时，你可以给它们应用额外的**类型修饰符**来影响这些值的验证。我们先来看一个例子：

```graphql
type Character {
  name: String!
  appearsIn: [Episode]!
}
```

非空和列表修饰符可以组合使用。例如你可以要求一个非空字符串的数组：

```graphql
myField: [String!]
```

这表示**数组本身**可以为空，但是其不能有任何空值成员。用 JSON 举例如下：

```js
myField: null // 有效
myField: [] // 有效
myField: ['a', 'b'] // 有效
myField: ['a', null, 'b'] // 错误
```

然后，我们来定义一个不可为空的字符串数组：

```graphql
myField: [String]!
```

这表示数组本身不能为空，但是其可以包含空值成员：

```js
myField: null // 错误
myField: [] // 有效
myField: ['a', 'b'] // 有效
myField: ['a', null, 'b'] // 有效
```

你可以根据需求嵌套任意层非空和列表修饰符。

### 6. 接口类型

例如，你可以用一个 `Character` 接口用以表示《星球大战》三部曲中的任何角色：

```graphql
interface Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
}
```

这意味着任何**实现** `Character` 的类型都要具有这些字段，并有对应参数和返回类型。

例如，这里有一些可能实现了 `Character` 的类型：

```graphql
type Human implements Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
  starships: [Starship]
  totalCredits: Int
}

type Droid implements Character {
  id: ID!
  name: String!
  friends: [Character]
  appearsIn: [Episode]!
  primaryFunction: String
}
```

### 7. 联合类型

联合类型和接口十分相似，但是它并不指定类型之间的任何共同字段。

```graphql
union SearchResult = Human | Droid | Starship
```

在我们的schema中，任何返回一个 `SearchResult` 类型的地方，都可能得到一个 `Human`、`Droid` 或者 `Starship`。注意，联合类型的成员需要是具体对象类型；你不能使用接口或者其他联合类型来创造一个联合类型。

这时候，如果你需要查询一个返回 `SearchResult` 联合类型的字段，那么你得使用条件片段才能查询任意字段。

```
{
  search(text: "an") {
    ... on Human {
      name
      height
    }
    ... on Droid {
      name
      primaryFunction
    }
    ... on Starship {
      name
      length
    }
  }
}
```

```
{
  "data": {
    "search": [
      {
        "name": "Han Solo",
        "height": 1.8
      },
      {
        "name": "Leia Organa",
        "height": 1.5
      },
      {
        "name": "TIE Advanced x1",
        "length": 9.2
      }
    ]
  }
}
```



### 8. 输入类型

目前为止，我们只讨论过将例如枚举和字符串等标量值作为参数传递给字段，但是你也能很容易地传递复杂对象。这在变更（mutation）中特别有用，因为有时候你需要传递一整个对象作为新建对象。在 GraphQL schema language 中，输入对象看上去和常规对象一模一样，除了关键字是 `input` 而不是 `type`：

```graphql
input ReviewInput {
  stars: Int!
  commentary: String
}
```

输入对象类型上的字段本身也可以指代输入对象类型，但是你不能在你的 schema 混淆输入和输出类型。输入对象类型的字段当然也不能拥有参数。





# http实践

raphQL 服务器在单个 URL /入口端点（通常是 `/graphql`）上运行，并且所有提供服务的 GraphQL 请求都应被导向此入口端点。



你的 GraphQL HTTP 服务器应当能够处理 HTTP GET 和 POST 方法。

### GET 请求 

在收到一个 HTTP GET 请求时，应当在 “query” 查询字符串（query string）中指定 GraphQL 查询。例如，如果我们要执行以下 GraphQL 查询：

```graphql
{
  me {
    name
  }
}
```

此请求可以通过 HTTP GET 发送，如下所示：

```undefined
http://myapi/graphql?query={me{name}}
```

查询变量可以作为 JSON 编码的字符串发送到名为 `variables` 的附加查询参数中。如果查询包含多个具名操作，则可以使用一个 `operationName` 查询参数来控制哪一个应当执行。

### POST 请求 

标准的 GraphQL POST 请求应当使用 `application/json` 内容类型（content type），并包含以下形式 JSON 编码的请求体：

```js
{
  "query": "...",
  "operationName": "...",
  "variables": { "myVariable": "someValue", ... }
}
```

`operationName` 和 `variables` 是可选字段。仅当查询中存在多个操作时才需要 `operationName`。

除了上边这种请求之外，我们还建议支持另外两种情况：

- 如果存在 “query” 这一查询字符串参数（如上面的 GET 示例中），则应当以与 HTTP GET 相同的方式进行解析和处理。
- 如果存在 “application/graphql” Content-Type 头，则将 HTTP POST 请求体内容视为 GraphQL 查询字符串。



# 分页

我们有很多种方法来实现分页：

- 我们可以像这样 `friends(first:2 offset:2)` 来请求列表中接下来的两个结果。
- 我们可以像这样 `friends(first:2 after:$friendId)`, 来请求我们上一次获取到的最后一个朋友之后的两个结果。
- 我们可以像这样 `friends(first:2 after:$friendCursor)`, 从最后一项中获取一个游标并使用它来分页。

为了解决这两个问题，我们的 `friends` 字段可以返回一个连接对象。然后，连接对象将具有边其中的字段以及其他信息（如总计数和有关下一页是否存在的信息）。所以我们的最终查询可能看起来像这样：

```graphql
{
  hero {
    name
    friends(first:2) {
      totalCount
      edges {
        node {
          name
        }
        cursor
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
```







# graphql-go 的实现demo

```go
package main

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

// 定义good object
type Goods struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Url   string  `json:"url"`
}

var goodsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Goods",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var goodsListType = graphql.NewList(goodsType)

// 定义query
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"goodsList": &graphql.Field{
				Type: goodsListType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return []Goods{}, nil
				},
			},
			"goods": &graphql.Field{
				Type: goodsType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return Goods{
							ID:    idQuery,
							Name:  "test",
							Price: 222.222,
							Url:   "https://jiangshiyi.top",
						}, nil
					}
					err := errors.New("Field 'goods' is missing required arguments: id. ")
					return nil, err
				},
			},
		},
	})

// 定义input类型
var goodsInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "goodsInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"url": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})

// 定义mutation
var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addGoods": &graphql.Field{
				Type: goodsType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: goodsInputType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input, isOK := p.Args["input"].(map[string]interface{})
					if !isOK {
						err := errors.New("Field 'addGoods' is missing required argument input. ")
						return nil, err
					}
					result := Goods{
						Name:  input["name"].(string),
						Price: input["price"].(float64),
						Url:   input["url"].(string),
					}
					return result, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)

}

// 第三方库 的使用
/*func ginHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()
	router.Any("/graphql", ginHandler())
	router.Run(":8080")
}*/

```

