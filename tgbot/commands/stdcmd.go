package commands

var StandardCommandHandlers = []Command{
	&EchoCommand{},
	&HelpCommand{},
	&NewMemberCommand{},
	&TimeCommand{},
	&StatsCommand{},
	&LeetCommand{},
	&MentionedCommand{},
	//&GlobalCommand{},
}
