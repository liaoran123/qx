package aiemot

import "fmt"

//情绪词
type Emot struct {
	emot  string //情绪词
	not   bool   //是否有否定
	level int    //程度词
}

// 情绪分析模型接口
type AnalysisEmotion interface {
	Analysis(emot []Emot)
}

// 情绪词典
type Dict struct {
}

//正负模型
type Minus struct {
	temp string
}

func (m *Minus) Analysis(emot []Emot) {
	fmt.Printf("emot: %v\n", emot)
}

func main() {
	minus := &Minus{}
	var ae AnalysisEmotion = minus
	ae.Analysis(nil)
}

/*
算法逻辑
Step 1：读取评论数据，对评论进行分句（分句主要以特定的标点符号为主）。
Step 2：将结巴词典和所有情感词典做并集，得出新的分词词典。
Step 3：查找分句的情感词，记录正面还是负面，以及位置。
Step 4：在情感词前查找程度词，找到就停止搜寻。为程度词设权值，乘以情感值。
Step 5：在情感词前查找否定词，找完全部否定词，若数量为奇数，乘以-1，若为偶数， 乘以 1。
Step 6：找出感叹号和问好等重要的标点符合
- 判断分句结尾是否有感叹号，有叹号则往前寻找情感词，有则相应的情感值+2。
- 判断分句结尾是否有问好，有问号该句判断为负面值+2。
Step 7：计算完一条评论所有分句的情感值（[正面分值, 负面分值]），用数组（list） 记录起来。
Step 8：计算每条评论中每一个分句的的正面情感均值与负面情感均值，然后比较正面情感总和与负面情感总和，较大的一个即为所得情感倾向。
*/
/*
func NewAnalysisEmotion(text string) *AnalysisEmotion {
	return &AnalysisEmotion{
		text: text,
	}
}
func (a *AnalysisEmotion) Aaa() {

}
*/
