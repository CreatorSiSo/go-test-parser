package preprocessor

import (
	"fmt"
	"testing"
)

const html = `<!DOCTYPE html>
<html class="no-js" lang="de">
  <head>
    {app:{head:{default}}}

    <title>SchillerGarten - Start</title>
  </head>

  <body>
    <div hidden>
      {app:{head:{icons}}}
      <!--  -->
      {app:{video-player:{video-controls}}}
    </div>

    {app:{sitemap}}
    <!--  -->
    {app:{image-viewer}}
    <!--  -->
    {app:{header}}
    <!--  -->
    {app:{backdrop}}

    <main>
      {app:{slider:{hero-slider}}}
      <!--  -->
      {app:{quick-start-nav}}
      <!--  -->
      {app:{content-section:{news-events}}}
      <!--  -->
      {app:{content-section:{ice}}}
      <!--  -->
      {app:{content-section:{news-events}}}
      <!--  -->
      {app:{content-section:{ice}}}
      {app:{content-section:{news-events}}}
      {app:{content-section:{ice}}}
      {app:{content-section:{news-events}}}
      {app:{content-section:{ice}}}
        -->
      {app:{content-section:{news-events}}}
      <!--  -->
      {app:{content-section:{ice}}}
      <!--  -->
      {app:{content-section:{activities}}}
    </main>

    {app:{footer}}
    <!--
    {app:{scroll-to-top}}
  </body>
</html>

<!DOCTYPE html>
<html lang="en">
  <head>-->
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    {sdlfhh:{}
    <div>
      :{sduhsfhui}
      <section>{app:{test:{test}}}</section>
      <section>{app:{test}}</section>
      <!--

		<section>{app:{comp:{}}}</section>

		<button></button>
		-->
    </div>
  </body>
</html>`

const html2 = `<!DOCTYPE html>
<html lang="en">
  <head>-->
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    {sdlfhh:{}
    <div>
      :{sduhsfhui}
      <section>{app:{test:{test}}}</section>
      <section>{app:{test}}</section>
      <!--

		<section>{app:{comp:{}}}</section>

		<button></button>
		-->
    </div>
  </body>
</html>`

var tokenizer = Tokenizer{
	spec: []TokenSpec{
		{
			regExp:     `^<!--[\s\S]*?(-->)`,
			matchIndex: 0,
			tokenType:  COMMENT,
		},
		{
			regExp:     `^{[\s\S]*?}+`,
			matchIndex: 0,
			tokenType:  EXPRESSION,
		},
		{
			regExp:     `^([\s\S]*?)($|<!--|{)`,
			matchIndex: 1,
			tokenType:  CONTENT,
		},
	},
}

func printTokens(tokens []Token) {
	fmt.Print("[\n")
	for _, token := range tokens {
		fmt.Printf("%s: \"%s\",\n", token.tokenType.String(), token.content)
	}
	fmt.Print("]\n")
}

func TestTokenizeHTML(test *testing.T) {
	tokenizer.input = html
	tokenizer.cursor = 0

	var tokens []Token

	for tokenizer.cursor < len(tokenizer.input) {
		nextToken := tokenizer.NextToken()

		if nextToken == nil {
			break
		}

		tokens = append(tokens, *nextToken)
	}

	printTokens(tokens)
}

func TestTokenizeHTML2(test *testing.T) {
	tokenizer.input = html2
	tokenizer.cursor = 0

	var tokens []Token

	for tokenizer.cursor < len(tokenizer.input) {
		nextToken := tokenizer.NextToken()

		if nextToken == nil {
			break
		}

		tokens = append(tokens, *nextToken)
	}

	printTokens(tokens)
}

// func TestTokenizeEmpty(test *testing.T) {
// 	tokenizer.input = ""
// 	tokenizer.cursor = 0

// 	nextToken := tokenizer.NextToken()
// 	nextToken = tokenizer.NextToken()
// 	nextToken = tokenizer.NextToken()

// 	fmt.Println(nextToken)
// }

// Token{
// 	content: "<!-- <section>{app:{test:{test}}}</section> -->",
// 	token:   COMMENT,
// },
// Token{
// 	content: "{svg:{ashg}}",
// 	token:   EXPRESSION,
// }
