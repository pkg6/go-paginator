package paginator

import (
	"errors"
	"math"
)

type IAdapter interface {
	// Length 数据长度
	Length() (int64, error)
	// Slice 切割数据(分页)
	Slice(offset, length int64, dest any) error
}
type IPaginator interface {
	Clear()
	// Get 获取数据
	Get(dest any) error
	// SetDestItems 设置输出源
	SetDestItems(dest any)
	// GetDestItems 获取输出源
	GetDestItems(dest any) any
	// GetResult 获取 Result结构体数据
	GetResult(dest any) *Result
	// SetCurrentPage 设置当前页数
	SetCurrentPage(currentPage int64)
	// GetCurrentPage 获取当前页页码
	GetCurrentPage() int64
	// GetTotal 获取数据总条数
	GetTotal() (int64, error)
	// GetListRows 获取每页数量
	GetListRows() int64
	//GetLastPage 获取最后一页页码
	GetLastPage() (int64, error)
	//HasPages 数据是否足够分页
	HasPages() bool
}

type Paginator struct {
	Simple bool
	//数据集适配器
	IAdapter IAdapter
	//设置输出源
	DestItems any
	//当前页
	CurrentPage int64
	//最后一页
	LastPage int64
	//数据总数
	Total int64
	//每页数量
	ListRows int64
	//是否有下一页
	HasMore bool
	//分页配置
	Options Options
}

func (p Paginator) Clear() {
	p.Simple = false
	p.IAdapter = nil
	p.DestItems = nil
	p.CurrentPage = 0
	p.LastPage = 0
	p.Total = 0
	p.ListRows = 0
	p.HasMore = false
	p.Options = DefaultOptions()

}
func SimplePaginator(adapter IAdapter, listRows, currentPage int64) IPaginator {
	return Make(adapter, listRows, currentPage, 0, true, DefaultOptions())
}
func TotalPaginator(adapter IAdapter, listRows, currentPage, total int64) IPaginator {
	return Make(adapter, listRows, currentPage, total, false, DefaultOptions())
}
func Make(adapter IAdapter, listRows int64, currentPage int64, total int64, simple bool, options Options) IPaginator {
	p := &Paginator{}
	p.Clear()
	p.Simple = simple
	p.IAdapter = adapter
	p.ListRows = listRows
	p.Options = options
	if p.Simple {
		p.SetCurrentPage(currentPage)
		count, err := p.IAdapter.Length()
		if err != nil {
			panic(err)
		}
		p.HasMore = count > p.CurrentPage
	} else {
		p.Total = total
		p.LastPage = int64(math.Ceil(float64(p.Total) / float64(p.ListRows)))
		p.SetCurrentPage(currentPage)
		p.HasMore = p.CurrentPage < p.LastPage
	}
	return p
}

func (p *Paginator) SetDestItems(dest any) {
	p.DestItems = dest
}
func (p *Paginator) GetDestItems(dest any) any {
	if dest != nil {
		p.DestItems = dest
	}
	if p.DestItems == nil {
		panic("DestItems uninitialized")
	}
	return p.DestItems
}

// SetCurrentPage 设置当前页数
func (p *Paginator) SetCurrentPage(currentPage int64) {
	if !p.Simple && currentPage > p.LastPage {
		if p.LastPage > 0 {
			p.CurrentPage = p.LastPage
			return
		} else {
			p.CurrentPage = 1
			return
		}
	}
	p.CurrentPage = currentPage
}

// GetTotal 获取数据总条数
func (p *Paginator) GetTotal() (int64, error) {
	if p.Simple {
		return 0, errors.New("not support total")
	}
	return p.Total, nil
}

// GetListRows 获取每页数量
func (p *Paginator) GetListRows() int64 {
	return p.ListRows
}

// GetCurrentPage 获取当前页页码
func (p *Paginator) GetCurrentPage() int64 {
	return p.CurrentPage
}

// GetLastPage 获取最后一页页码
func (p *Paginator) GetLastPage() (int64, error) {
	if p.Simple {
		return 0, errors.New("not support last")
	}
	return p.LastPage, nil
}

// HasPages 数据是否足够分页
func (p *Paginator) HasPages() bool {
	return !(p.CurrentPage == 1 && !p.HasMore)
}

// Get 获取指定长度的数据
func (p *Paginator) Get(dest any) error {
	var offset int64
	dest = p.GetDestItems(dest)
	page := p.GetCurrentPage()
	if page > 1 {
		offset = (page - 1) * p.ListRows
	}
	return p.IAdapter.Slice(offset, p.ListRows, dest)
}

// GetResult 获取 Result结构体数据
func (p *Paginator) GetResult(dest any) *Result {
	ret := &Result{}
	ret.Total, _ = p.GetTotal()
	_ = p.Get(dest)
	ret.Data = dest
	ret.PerPage = p.GetListRows()
	ret.LastPage, _ = p.GetLastPage()
	ret.CurrentPage = p.GetCurrentPage()
	return ret
}
