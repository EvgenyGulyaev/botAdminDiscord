package command

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var blackList = map[string]bool{
	"471650968245764099": true, // Гоша
	"240881052720037888": true, // Антон
}

func HandleChangeNick(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		showNickHelp(s, m)
		return
	}
	subCommand := strings.ToLower(args[0])
	switch subCommand {
	case "set", "change", "edit":
		handleNickSet(s, m, args[1:])
	default:
		if len(m.Mentions) > 0 {
			handleNickSet(s, m, args)
		} else {
			showNickHelp(s, m)
		}
	}
}

func showNickHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	sendMessage(s, m, "!nick set @User123 НовыйНик")
}

func handleNickSet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		sendMessage(s, m, "❌ Не указаны все аргументы.")
		return
	}

	if len(m.Mentions) == 0 {
		sendMessage(s, m, "❌ Упомяните пользователя, которому хотите изменить никнейм.")
		return
	}

	targetUser := m.Mentions[0]

	_, isBadUser := blackList[targetUser.ID]

	if isBadUser {
		sendMessage(s, m, "❌ У вас нет прав на смену ника.")
		return
	}

	nickStart := -1
	for i, arg := range args {
		if strings.HasPrefix(arg, "<@") && strings.HasSuffix(arg, ">") {
			nickStart = i + 1
			break
		}
	}

	if nickStart == -1 || nickStart >= len(args) {
		sendMessage(s, m, "❌ Укажите новый никнейм.")
		return
	}

	newNick := strings.Join(args[nickStart:], " ")

	err := s.GuildMemberNickname(m.GuildID, targetUser.ID, newNick)
	if err != nil {
		handleNickError(s, m, err)
		return
	}

	response := fmt.Sprintf("✅ **%s** изменил никнейм пользователя **%s**\n"+
		"**Новый никнейм:** %s",
		m.Author.Username,
		targetUser.Username,
		newNick)

	sendMessage(s, m, response)
}

func handleNickError(s *discordgo.Session, m *discordgo.MessageCreate, err error) {
	errMsg := err.Error()
	res := ""

	switch {
	case strings.Contains(errMsg, "Missing Permissions"):
		res = "❌ У бота нет прав на изменение никнеймов.\n" +
			"Дайте боту право **'Manage Nicknames'** в настройках сервера."

	case strings.Contains(errMsg, "rate limited"):
		res = "❌ Слишком много запросов. Discord ограничивает частоту изменений никнеймов."

	case strings.Contains(errMsg, "Cannot edit nickname of higher hierarchy"):
		res = "❌ Нельзя изменить никнейм пользователю с более высокой ролью."

	default:
		res = fmt.Sprintf("❌ Ошибка: %v", err)
	}
	sendMessage(s, m, res)
}
