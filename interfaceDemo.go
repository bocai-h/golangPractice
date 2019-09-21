package main

import "fmt"

//AnimalCategory 动物分类
type AnimalCategory struct {
	kingdom string //界
	phylum  string //门
	class   string //纲
	order   string //目
	family  string //科
	genus   string //属
	species string //种
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s-%s-%s", ac.kingdom, ac.phylum, ac.class, ac.order, ac.family, ac.genus, ac.species)
}

//Animal 动物
type Animal struct {
	scientificName string
	AnimalCategory
}

//Category 动物类别
func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func (a Animal) String() string {
	return fmt.Sprintf("%s category: %s", a.scientificName, a.AnimalCategory)
}

//Cat 猫类
type Cat struct {
	name string
	Animal
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %s)", cat.scientificName, cat.Animal, cat.name)
}

//SetName 设置名字
func (cat *Cat) SetName(name string) {
	cat.name = name
	return
}

//Name 返回名称
func (cat *Cat) Name() string {
	if cat != nil {
		return cat.name
	}
	return "blank name"
}

//Pet 宠物类型
type Pet interface {
	// SetName(name string)
	Name() string
	Category() string
}

func main() {
	// cat := Cat{name: "小猫"}
	// var pet Pet
	// pet = cat
	// fmt.Printf("pet.Name=%s\n", pet.Name())

	// cat.SetName("monster")
	// fmt.Printf("cat.Name=%s\n", cat.Name())
	// fmt.Printf("pet.Name=%s\n", pet.Name())

	var pet Pet
	var catptr *Cat
	pet = catptr

	if pet == nil {
		fmt.Printf("pet is nil")
	}

	fmt.Printf("catptr=%v\n", catptr)
	fmt.Printf("pet Type = %T\n", pet)
	fmt.Printf("pet name = %s\n", pet.Name())
}
