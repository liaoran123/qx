package db

import (
	"fmt"

	"github.com/wangbin/jiebago"
)

// 对应diary表
type Diary struct {
	Tid     int               `json:"tid"`     //表id
	Uid     int               `json:"Uid"`     //用户id
	Title   string            `json:"title"`   //标题
	Type    string            `json:"type"`    //情绪类型
	Content string            `json:"content"` //内容
	Dtime   string            `json:"dtime"`   //发布时间
	play    bool              `json:"-"`       //付款状态：0,未付款；1，付款。用于统计过滤
	seg     jiebago.Segmenter `json:"-"`
	Kws     []*Kws            `json:"es"`
}

func NewDiary() *Diary {
	return &Diary{
		play: false,
	}
}

// 初始jiebago化词库
func (d *Diary) Jbinit() {
	d.seg.LoadDictionary("dict/jieba_sentiment.txt")
}

// 获取用户付款状态
func (d *Diary) Getpaly() bool {
	return true
}

// 插入表
func (d *Diary) Add(param map[string]string) bool {
	return true
}

// 切词
func (d *Diary) Segmentation(text string) {
	d.Jbinit()
	/*
		fmt.Print("【全模式】：")
		print(seg.CutAll("我来到北京清华大学"))
		//【全模式】： 我 / 来到 / 北京 / 清华 / 清华大学 / 华大 / 大学 /

		fmt.Print("【精确模式】：")
		print(seg.Cut("我来到北京清华大学", false))
		//【精确模式】： 我 / 来到 / 北京 / 清华大学 /

		fmt.Print("【新词识别】：")
		print(seg.Cut("他来到了网易杭研大厦", true))
		//【新词识别】： 他 / 来到 / 了 / 网易 / 杭研 / 大厦 /

		fmt.Print("【搜索引擎模式】：")
		print(seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))
		//【搜索引擎模式】： 小明 / 硕士 / 毕业 / 于 / 中国 / 科学 / 学院 / 科学院 / 中国科学院 / 计算 / 计算所 / ， / 后 / 在 / 日本 / 京都 / 大学 / 日本京都大学 / 深造 /
	*/
	kws := d.seg.Cut(text, false)
	fmt.Printf("kws: %v\n", kws)
	for k := range kws {
		d.Kws = append(d.Kws, NewKws(k))
	}
	fmt.Printf("kws: %v\n", kws)
	//查找词在text中的位置
	for _, e := range d.Kws {
		fmt.Printf("e: %v\n", e)
	}
}

// 修改
func (d *Diary) Udate(param map[string]string) bool {
	return true
}

// 删除
func (d *Diary) Del(param map[string]string) bool {
	//删除Kws子表相关所有记录
	Kws := new(Kws)
	Kws.Dels(111)
	//删除tid的记录
	return true
}
