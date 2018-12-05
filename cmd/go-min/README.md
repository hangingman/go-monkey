# Minの文法

* BNF記法による言語仕様
```
<プログラム>::=[<変数宣言>]<文>
<変数宣言>::=var<変数名>{","<変数名>}";"
<文字>::=a|b|...|z
<変数名>::=<文字>{<文字>}{<数字>}
<正数字>::=1|2|...|9
<数字>::=0|<正数字>
<定数>::=0|<正数字>{<数字>}
<式>::=<変数名>|<定数>|<式>+<式>|<式>*<式>|<式>-<式>|<式>"="<式>|<式>>=<式>|"("<式>")"
<文>::=<単純文>|<複合文>
<単純文>::=input<変数名>|output<変数名>|<変数名>:=<式>|if<式>then<文>[else<文>]fi|while<式>begin<文>end
<複合文>::=<単純文>";"<単純文>{";"<単純文>}
```

* プログラム用に英文表記
```
<Program>::=[<VariableDeclaration>]<Statement>
<VariableDeclaration>::=var<VariableName>{","<VariableName>}";"
<Char>::=a|b|...|z
<VariableName>::=<Char>{<Char>}{<Number>}
<NaturalNumber>::=1|2|...|9
<Number>::=0|<NaturalNumber>
<Constant>::=0|<NaturalNumber>{<Number>}
<Expression>::=<VariableName>|<Constant>|<Expression>+<Expression>|<Expression>*<Expression>|<Expression>-<Expression>|<Expression>"="<Expression>|<Expression>>=<Expression>|"("<Expression>")"
<Statement>::=<SimpleStatement>|<CompoundStatement>
<SimpleStatement>::=input<VariableName>|output<VariableName>|<VariableName>:=<Expression>|if<Expression>then<Statement>[else<Statement>]fi|while<Expression>begin<Statement>end
<CompoundStatement>::=<SimpleStatement>";"<SimpleStatement>{";"<SimpleStatement>}
```
