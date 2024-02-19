package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var pokers []int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入4个扑克牌数字，A,J,Q,K 用1,11,12,13 代替，以空格分隔：")
	// 读取一行用户输入
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // 移除前后空白字符

	// 按空格分割字符串为子串
	numStrings := strings.Split(input, " ")
	if len(numStrings) != 4 { // 检查输入的数量是否正确
		fmt.Println("输入的数量不正确，请确保输入了4个数字")
		return
	}

	// 将子串转换为整数并存入数组
	for _, str := range numStrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("无法将", str, "转换为数字:", err)
			return
		}
		pokers = append(pokers, num)
	}

	// 输出数组内容验证结果
	fmt.Println("您输入的4张牌是：", pokers)

	// 1. 洗牌，随意打乱
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(pokers), func(i, j int) {
		pokers[i], pokers[j] = pokers[j], pokers[i]
	})
	fmt.Printf("1. 随机洗牌后的牌：%v\n", pokers)

	// 2. 对折，然后撕开，即切片内容 * 2
	pokers = append(pokers, pokers...)
	fmt.Printf("2. 对折后,然后撕开的牌：%v\n", pokers)

	// 3. 问问自己名字有几个字，就从最上面拿出对应个数的牌放到底部，例如「刘谦」名字有 2 个字，即将切片前 2 个元素取出，放到切片最后
	fmt.Println("3.请问你的名字有几个字：")
	// 读取一行用户输入
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input) // 移除前后空白字符
	nameLen, _ := strconv.Atoi(input)
	pokers = append(pokers[nameLen:], pokers[:nameLen]...)
	fmt.Printf("3.1 你的名字有%d个字，就从最上面拿出对应个数的牌放到底部：%v\n", nameLen, pokers)

	// 4. 拿起最上面的 3 张牌，插入中间任意位置，这里将切片前 3 个元素取出，并将其插入最后一个元素之前
	pokers = append(pokers[3:7], pokers[0], pokers[1], pokers[2], pokers[7])
	fmt.Printf("4. 拿起最上面的 3 张牌，插入中间任意位置: %v\n", pokers)

	// 5. 拿出最上面的 1 张牌，藏于秘密的地方，比如屁股下，这里使用 top 变量暂存
	top := pokers[0]
	pokers = pokers[1:]
	fmt.Printf("5. 拿出最上面的 1 张牌，藏于秘密的地方：%d, %v\n", top, pokers)

	// 6. 如果你是南方人，从上面拿起 1 张牌；如果你是北方人，则从上面拿起 2 张牌；假如我们不确定自己是南方人还是北方人，那就干脆拿起 3 张牌，然后插入中间任意位置
	// 读取一行用户输入
	fmt.Println("6. 如果你是南方人还是北方人？ 南方人输入1，北方人输入2，不确定输入3：")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input) // 移除前后空白字符
	southornorth, _ := strconv.Atoi(input)
	if southornorth == 1 {
		pokers = append(pokers[1:6], pokers[0], pokers[6])
		fmt.Printf("6.1 你是南方人,从上面拿出 1 张牌: %v\n", pokers)
	} else if southornorth == 2 {
		pokers = append(pokers[2:7], pokers[0], pokers[1], pokers[7])
		fmt.Printf("6.1 你是北方人,从上面拿出 3 张牌: %v\n", pokers)
	} else {
		pokers = append(pokers[3:7], pokers[0], pokers[1], pokers[2], pokers[7])
		fmt.Printf("6.1 不确定南方还是北方人，从上面拿出 3 张牌: %v\n", pokers)
	}
	fmt.Println("7. 请问性别是？男生输入1，女生请输入2：")
	// 读取一行用户输入
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input) // 移除前后空白字符
	sex, _ := strconv.Atoi(input)
	if sex == 1 {
		// 7. 如果你是男生，从上面拿起 1 张牌；如果你是女生，则从上面拿起 2 张牌，撒到空中（扔掉）
		pokers = pokers[1:]
		fmt.Printf("7.1 你是男生，从上面拿起 1 张牌，撒到空中（扔掉）：%v\n", pokers)
	} else if sex == 2 {
		// 7. 如果你是女生，从上面拿起 2 张牌；如果你是男生，则从上面拿起 1 张牌，撒到空中（扔掉）
		pokers = pokers[2:]
		fmt.Printf("7.1 你是女生，从上面拿起 2 张牌，撒到空中（扔掉）：%v\n", pokers)
	} else {
		fmt.Println("输入错误！")
		return
	}
	// 8. 魔法时刻，在遥远的魔术的历史上，流传了一个七字真言「见证奇迹的时刻」，可以带给我们幸福。现在，我们每念一个字，从上面拿一张放到最底部，即需要完成 7 次同样的操作
	//    我们可以用一个 `for loop` 实现
	fmt.Printf("8. 魔法时刻，在遥远的魔术的历史上，流传了一个七字真言「见证奇迹的时刻」，可以带给我们幸福。现在，我们每念一个字，从上面拿一张放到最底部，即需要完成 7 次同样的操作\n")

	for k, v := range []string{"见", "证", "奇", "迹", "的", "时", "刻"} {
		pokers = append(pokers[1:], pokers[0])
		fmt.Printf("8.%d %s：%v\n", k+1, v, pokers)

	}
	fmt.Printf("9. 最后一个环节，叫「好运留下来，烦恼丢出去」，在念到「好运留下来」时，从上面拿起 1 张牌放入底部；在念到「烦恼丢出去」时，从上面拿起 1 张牌扔掉，女生需要完成 4 次同样的操作，男生需要完成 5 次同样的操作\n")

	// 9. 最后一个环节，叫「好运留下来，烦恼丢出去」，在念到「好运留下来」时，从上面拿起 1 张牌放入底部；在念到「烦恼丢出去」时，从上面拿起 1 张牌扔掉，女生需要完成 4 次同样的操作，男生需要完成 5 次同样的操作
	if sex == 1 {
		// 7. 如果你是男生，从上面拿起 1 张牌；如果你是女生，则从上面拿起 2 张牌，撒到空中（扔掉）
		for _, v := range []int{1, 2, 3, 4, 5} {
			// 好运留下来
			pokers = append(pokers[1:], pokers[0])
			// 烦恼丢出去
			pokers = pokers[1:]
			fmt.Printf("9.%d 好运留下来，烦恼丢出去 ：%v\n", v, pokers)
		}
	} else if sex == 2 {
		for _, v := range []int{1, 2, 3, 4} {
			// 好运留下来
			pokers = append(pokers[1:], pokers[0])
			// 烦恼丢出去
			pokers = pokers[1:]
			fmt.Printf("9.%d 好运留下来，烦恼丢出去 ：%v\n", v, pokers)
		}
	} else {
		fmt.Println("输入错误！")
		return
	}
	// 最后，我们将见证奇迹：
	fmt.Printf("见证奇迹：%d == %d", top, pokers[0])
}
