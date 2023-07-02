{{/*    
    Automatically adds role categories to roles 
        Trigger: Added + Removed reactions 
    Trigger Type:  Reaction
    Only run in the following channels: Where the roles are set 
    
    Copyright: 2023 Grafik#9575
    License: Pay me.
*/}}
 
{{/*settings */}}
{{$category  := `- - -`}}
 
{{/* Don't touch */}}

{{$userID := .Reaction.UserID}}
{{$guildRoles := .Guild.Roles}}
{{$userRoles := (getMember $userID).Roles}}

{{$categories := dict}}
{{$sortedRoles := cslice}}

{{- range $guildRoles}}
    {{- if ne .Position 0}}
        {{- $sortedRoles = $sortedRoles.Append (sdict "position" .Position "name" .Name "id" .ID)}}
    {{- end}}
{{- end}}

{{$sortedRoles = sort $sortedRoles (sdict "key" "position" "reverse" true)}}

{{$currentCategory := 0}}

{{- range $sortedRoles}}
    {{- if (reFind $category .name)}}
        {{- $currentCategory = .id}}
        {{- $categories.Set .id cslice}}
    {{- else if $currentCategory}}
        {{- $currentRoles := $categories.Get $currentCategory}}
        {{- $currentRoles = $currentRoles.Append .id}}
        {{- $categories.Set $currentCategory $currentRoles}}
    {{- else}}
        {{/* ROLE HAS NO CATEGORY */}}
    {{- end}}
{{- end}}

{{- range $category, $roles := $categories}}
    {{- $isEmpty := true}}
    {{- range $roles}}
        {{- if in $userRoles .}}
            {{- $isEmpty = false}}
            {{- if not (in $userRoles $category)}}
                {{- addRoleID $category}}
            {{- end}}
        {{- end}}
    {{- end}}

    {{- if and $isEmpty (in $userRoles $category)}} 
        {{- removeRoleID $category}}
    {{- end}}
{{- end}}