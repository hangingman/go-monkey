# RCFの文法

* BNF記法による言語仕様
```
<プログラム>::=[<宣言部>]<本体>
<宣言部>::=[<大域変数宣言部>][<関数宣言部>]
<大域変数宣言部>::=var<変数列>";"
<関数宣言部>::=def<関数宣言>in
<関数宣言>::=<関数定義式>{";"<関数定義式>}
<関数定義式>::=<関数名>"("<局所変数宣言>")""="<式>
<局所変数宣言>::=<変数列>
<変数列>::=[<変数名>{","<変数名>}]
<本体>::=<式>
<式>::=<変数名>|<定数>|<式>+<式>|<式>*<式>|<式>-<式>|<式>>=<式>|"("<式>")"|<関数名>"("[<式>{","<式>}]")"|if<式>then<式>else<式>fi
<文字>::=a|b|...|z
<識別子>::=<文字>{<文字>|<数字>|}
<変数名>::=<識別子>
<関数名>::=<識別子>
<正数字>::=1|2|...|9
<数字>::=0|<正数字>
<定数>::=0|<正数字>{<数字>}
```

* プログラム用に英文表記
```
<Program>::=[<DeclarationPart>]<MainPart>                                                           
<DeclarationPart>::=[<GlobalVariableDeclarationPart>][<FunctionDeclarationPart>]                    
<GlobalVariableDeclarationPart>::=var<VariableArray>";"                                             
<FunctionDeclarationPart>::=def<FunctionDeclaration>in                                              
<FunctionDeclaration>::=<FunctionDefineExpression>{";"<FunctionDefineExpression>}                   
<FunctionDefineExpression>::=<FunctionName>"("<LocalVariableDeclaration>")""="<Expression>          
<LocalVariableDeclaration>::=<VariableArray>                                                        
<VariableArray>::=[<VariableName>{","<VariableName>}]                                               
<MainPart>::=<Expression>                                                                           
<Expression>::=<VariableName>|<Constant>|<Expression>+<Expression>|<Expression>*<Expression>|<Expre\
  ssion>-<Expression>|<Expression>>=<Expression>|"("<Expression>")"|<FunctionName>"("[<Expression>{",\
  "<Expression>}]")"|if<Expression>then<Expression>else<Expression>fi                                 
<Char>::=a|b|...|z                                                                                  
<Ident>::=<Char>{<Char>|<Number>|}                                                                  
<VariableName>::=<Ident>                                                                            
<FunctionName>::=<Ident>                                                                            
<NaturalNumber>::=1|2|...|9                                                                         
<Number>::=0|<NaturalNumber>                                                                        
<Constant>::=0|<NaturalNumber>{<Number>}
```
