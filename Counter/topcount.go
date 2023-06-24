{{$now := currentTime}}
{{$page := 1}}{{$id := 0}}
{{if .CmdArgs}}
{{$page = toInt (index .CmdArgs 0)}}
 
{{if lt $page 1}} {{$page = 1}} {{end}} {{end}} 
 
{{with .ExecData}}
{{$page = .page}} {{$id = .id}}
{{end}}
{{$num_per_page := 10}}
{{$skip := mult (sub $page 1) $num_per_page}}
 
 
 
{{$text := ""}} {{$count := add $skip 1}}
{{range (dbTopEntries `counter\_tracker\_%` 10 $skip)}}
{{$member := getMember (toInt (index (split .Key "_") 2))}}
{{if $member}}
{{$text = joinStr "\n" $text (printf "#%-4d %5d - %-20s" $count (toInt .Value) $member.User)}}
{{else}}
{{$text = joinStr "\n" $text (printf "#%-4d %5d - %-20d" $count (toInt .Value) (toInt (index (split .Key "_") 2)))}}
{{end}}
{{$count = add $count 1}}
{{end}}
 
 
{{if $text}}
{{$embed := cembed  "title" "Counter Leaderboard" "description" (print "``" "`# -- Points -- User\n" $text "``" "`") "footer" (sdict "text" (print "React with 🗑️ to delete this message.\nPage: " $page)) "color" (randInt 0 16777216) }}
 
{{with $id}}
{{editMessage nil $id $embed}}
{{else}}
{{addMessageReactions nil (sendMessageRetID nil $embed) "⬅️"  "➡️" "🗑" }}
{{end}}
 
{{end}}
 
{{/* {{currentTime.Sub $now}} */}}