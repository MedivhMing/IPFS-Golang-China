
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package commands

import (
	"strings"
	"testing"

	cmds "gx/ipfs/QmWGm4AbZEbnmdgVTza52MSNpEmBdFVqzmAysRbjrRyGbH/go-ipfs-cmds"
)

func checkHelptextRecursive(t *testing.T, name []string, c *cmds.Command) {
	c.ProcessHelp()

	t.Run(strings.Join(name, "_"), func(t *testing.T) {
		if c.External {
			t.Skip("external")
		}

		t.Run("tagline", func(t *testing.T) {
			if c.Helptext.Tagline == "" {
				t.Error("no Tagline!")
			}
		})

		t.Run("longDescription", func(t *testing.T) {
			t.Skip("not everywhere yet")
			if c.Helptext.LongDescription == "" {
				t.Error("no LongDescription!")
			}
		})

		t.Run("shortDescription", func(t *testing.T) {
			t.Skip("not everywhere yet")
			if c.Helptext.ShortDescription == "" {
				t.Error("no ShortDescription!")
			}
		})

		t.Run("synopsis", func(t *testing.T) {
			t.Skip("autogenerated in go-ipfs-cmds")
			if c.Helptext.Synopsis == "" {
				t.Error("no Synopsis!")
			}
		})
	})

	for subname, sub := range c.Subcommands {
		checkHelptextRecursive(t, append(name, subname), sub)
	}
}

func TestHelptexts(t *testing.T) {
	Root.ProcessHelp()
	checkHelptextRecursive(t, []string{"ipfs"}, Root)
}
