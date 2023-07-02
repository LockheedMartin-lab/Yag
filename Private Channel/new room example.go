{{/* Testline */}}
{{print "Start"}}


{{/*configuration area*/}}
{{$ChannelID := }}
{{$Channel_Role := 805112942898053161}}
{{$heading := 805115909773328425}}


{{/* Do not change below */}}
{{$UserID := .User.Mention}}
{{$Channel := .Channel.Name($ChannelID) }}

{{/* Testline */}}
{{print "Test 1"}}

{{$args := parseArgs 1 "-[channel name] [@username]"
	(carg "user" "user")}}

	
	

{{if $UserID hasRoleID ($channel_role)}}
	{{takeRoleID ($args.Get 0).ID $Channel_Role}}

	{{print ($args.Get 0).Mention " has no longer access to the privat channel :key:" $Channel ". "}}

	{{/* Testline */}}
	{{print "Test 2"}}


{{else}}

	{{/* Testline */}}
	{{print "Test 3"}}

	{{giveRoleID ($args.Get 0).ID $Channel_Role}}

 
	{{deleteTrigger 0}}
	{{print ($args.Get 0).Mention " received access to the privat channel :key:" $Channel ". "}}

{{end}}

{{/* Testline */}}
{{print "Test 4"}}




{{/* - - - - - - - - - - - - Start Title cc - - - - - - - - - - - - 
    Copyright: 2021 Mandara_Greenblade#3457*/}}


{/*settings */}}
{{$Category_identifiers  := `- - -`}}
{{/*⬆️settings ⬆️*/}}
 
{{/* DO NOT TOUCH BELOW UNLESS YOU KNOW WHAT YOU ARE DOING */}}
 
{{/* Initialize Some Variables */}}
{{$UserID := .User.Mention}}
{{$Guild_Roles := .Guild.Roles}}
{{$Roles := sdict "0" "0" }}
{{$Categories_Roles:= sdict "0" "0" }}
{{$Categories_Roles_Position := sdict "0" "0"}} 
 
 
 
{{range .Guild.Roles }}{{$Roles.Set  (toString .ID) "0"}}{{ if  (reFind  $Category_identifiers  ( .Name )) }}{{$Categories_Roles.Set (toString .ID) "0"}}{{$Categories_Roles_Position.Set (toString .ID) (toString .Position)}}{{ end }}{{ end }} {{$Roles.Del "0"}} {{$Categories_Roles.Del "0"}} {{$Categories_Roles_Position.Del "0"}}
 
{{range $struct, $v := $Categories_Roles_Position}}{{range $Guild_Roles  }}{{if  and  (lt .Position (toInt $v))  (eq ($Roles.Get (toString .ID))  ($Roles.Get (toString $struct))) (ne  .Name "@everyone")}}{{$Roles.Set (toString .ID) (toString $struct)}}{{end}} {{end}}{{end}}
 
{{range $struct, $v := $Categories_Roles}}{{$Roles.Del $struct}}{{end}}
 
{{range $struct, $v := $Roles }}{{if (targetHasRoleID $UserID $struct)}}{{$Categories_Roles.Set (toString $v) "1"}}{{end}}{{end}}
 
{{ range $struct, $v := $Categories_Roles}}
{{ if eq $v  "1"}} {{if not (targetHasRoleID $UserID $struct)}}{{giveRoleID $UserID  $struct}} {{ end }} 
{{else}} {{if (targetHasRoleID $UserID $struct)}} {{takeRoleID $UserID  $struct}}{{ end }}{{ end }}{{ end }} 


{{/* - - - - - - - - - - - - End Title cc - - - - - - - - - - - - */}}


{{/* Testline */}}
{{print "Test 5"}}


{{sendDM "Test"}}