```
package "github.com/pipe4/lang/pipe4/gogll"


GoGLL: ImportBlock;

ImportBlock: "import" "(" ImportStatements ")";

ImportStatements: Import | Import ImportStatements;

Import: ImportName ImportPath;
ImportName: localDeclarationName;
ImportPath: string;

string: '"' <not "\\\"" | '\\' any "\\\"nrt"> '"';
localDeclarationName: letter {letter|number|'-'|'_'} letter;
```
