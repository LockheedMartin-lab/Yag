{{/* start of some random ass customization */}}
{{$Channel := "dungeon"}}
{{$Channel_Role := 805112942898053161}}
{{$heading := 805115909773328425}}

{{/*end of the fcking custom*/}}
{{$args := parseArgs 1 "-[Channel name] [@Username]"
    (carg "user" "user")}}
 

{{if hasroleID ($args.Get 0).ID  $Channel_Role}}
	{{removeRoleID ($args.Get 0).ID $Channel_Role}}
{{else}}
	{{giveRoleID ($args.Get 0).ID $Channel_Role}}
	{{giveRoleID ($args.Get 0).ID $heading}}
{{end}}


{{deleteTrigger 0}}
{{print ($args.Get 0).Mention " has received access to :key:" $Channel}}



=========

{{sleep 1}}
 
{{if (hasRoleID 798175665924079666)}}
{{removeRoleID 798175665924079666}}
{{addReactions ":x:"}}
{{else}}
{{addRoleID  798175665924079666}}
{{addReactions ":white_check_mark:"}}
{{end}}


============

{{/* start of some random ass customization */}}
{{$Channel := "dungeon"}}
{{$Channel_Role := 805112942898053161}}
{{$heading := 805115909773328425}}

{{/*end of the fcking custom*/}}
{{$args := parseArgs 1 "-[Channel name] [@Username]"
    (carg "user" "user")}}
 
 
{{giveRoleID ($args.Get 0).ID $Channel_Role}}
{{giveRoleID ($args.Get 0).ID $heading}}

{{deleteTrigger 0}}
{{print ($args.Get 0).Mention " has received access to :key:" $Channel}}
