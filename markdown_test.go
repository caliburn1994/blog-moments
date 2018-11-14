package main

import (
	"fmt"
	"testing"
)

func Test_AddMarkdown_1(t *testing.T) {
	input := "我和我妈说，我哥若是有了孩子后，你的生活更加累。我妈回我，若是我哥不要孩子，那以后她看到别人有孙子就会很寂寞。\n\n" +
		"让我有种感觉 老一辈的人，很少能离开社会寻找真正的快乐"

	fmt.Println(AddMarkdown(input))
}
