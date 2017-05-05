package com.netease.game.protobuf;

/**
 * Created by genevent.
 */
// {{.Comment}}
public class {{.Name}}
{
{{range .Defs}} 
	// {{.Comment}}
	public final static int {{.Name}} = {{.Value}};
{{end}}
};
