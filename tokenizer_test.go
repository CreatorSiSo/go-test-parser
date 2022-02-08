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

var testTokenizer = Tokenizer{
	Spec: []TokenSpec{
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

func TestTokenizeHTML(test *testing.T) {
	for i, token := range testTokenizer.Tokenize(html) {
		fmt.Println(i, token.String())
	}
}

func TestTokenizeHTML2(test *testing.T) {
	for i, token := range testTokenizer.Tokenize(html2) {
		fmt.Println(i, token.String())
	}
}

func TestTokenizeEmpty(test *testing.T) {
	fmt.Println(testTokenizer.Tokenize(""))
}

func TestTokenTypeToString(test *testing.T) {
	if COMMENT.String() != "COMMENT" ||
		EXPRESSION.String() != "EXPRESSION" ||
		CONTENT.String() != "CONTENT" {
		test.Fail()
	}
}
