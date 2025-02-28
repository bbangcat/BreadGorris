package src

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const FUN_TEXT = "> `avatar [id or mention]`"

func Avatar(s *discordgo.Session, m *discordgo.MessageCreate, args []string) CmdResult {
	id := strings.Replace(strings.Replace(strings.Replace(args[2], "!", "", -1), "<@", "", -1), ">", "", -1)
	u, user_errs := s.User(id)
	if user_errs != nil {
		return CmdResult{"avatar", errors.New(fmt.Sprint("유저 가져오기 오류: ", user_errs.Error()))}
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprint(u.Username, u.Discriminator, " 님의 아바타입니다."),
		Color:  WHITE,
		Footer: &discordgo.MessageEmbedFooter{Text: m.Author.Username, IconURL: m.Author.AvatarURL("1024")},
		Image:  &discordgo.MessageEmbedImage{URL: u.AvatarURL("1024")},
	})
	return CmdResult{"avatar", nil}
}
