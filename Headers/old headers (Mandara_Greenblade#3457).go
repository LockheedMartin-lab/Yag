
{{/*    
    Automatically adds role categories to roles 
        Trigger: Added + Removed reactions 
    Trigger Type:  Reaction
    Only run in the following channels: Where the roles are set 
    
    Copyright: 2021 Mandara_Greenblade#3457
    License: Please ask Mandara_Greenblade#3457 before redistributing  (Bitte fragen Sie Mandara_Greenblade, bevor Sie es weitergeben )
*/}}
 
{{/*settings */}}
{{$Category_identifiers  := `- - -`}}
{{/*⬆️settings ⬆️*/}}
 
{{/* DO NOT TOUCH BELOW UNLESS YOU KNOW WHAT YOU ARE DOING */}}
 
{{/* Initialize Some Variables */}}
{{$UserID := .Reaction.UserID}}
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