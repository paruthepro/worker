package settings

import (
	"github.com/TicketsBot/common/permission"
	"github.com/TicketsBot/worker/bot/command"
	"github.com/TicketsBot/worker/bot/command/registry"
	"github.com/TicketsBot/worker/bot/dbclient"
	"github.com/TicketsBot/worker/bot/utils"
	"github.com/TicketsBot/worker/i18n"
	"github.com/rxdn/gdl/objects/interaction"
)

type AutoCloseExcludeCommand struct {
}

func (AutoCloseExcludeCommand) Properties() registry.Properties {
	return registry.Properties{
		Name:             "exclude",
		Description:      i18n.HelpAutoCloseExclude,
		Type:             interaction.ApplicationCommandTypeChatInput,
		PermissionLevel:  permission.Support,
		Category:         command.Settings,
		DefaultEphemeral: true,
	}
}

func (c AutoCloseExcludeCommand) GetExecutor() interface{} {
	return c.Execute
}

func (AutoCloseExcludeCommand) Execute(ctx registry.CommandContext) {
	ticket, err := dbclient.Client.Tickets.GetByChannel(ctx.ChannelId())
	if err != nil {
		ctx.HandleError(err)
		return
	}

	if ticket.Id == 0 {
		ctx.Reply(utils.Red, "Error", i18n.MessageNotATicketChannel)
		return
	}

	if err := dbclient.Client.AutoCloseExclude.Exclude(ctx.GuildId(), ticket.Id); err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.Reply(utils.Green, "Autoclose", i18n.MessageAutoCloseExclude)
}
