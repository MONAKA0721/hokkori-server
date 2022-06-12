// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/MONAKA0721/hokkori/ent/post"
	"github.com/MONAKA0721/hokkori/ent/predicate"
	"github.com/MONAKA0721/hokkori/ent/user"
)

// PostWhereInput represents a where input for filtering Post queries.
type PostWhereInput struct {
	Predicates []predicate.Post  `json:"-"`
	Not        *PostWhereInput   `json:"not,omitempty"`
	Or         []*PostWhereInput `json:"or,omitempty"`
	And        []*PostWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "create_time" field predicates.
	CreateTime      *time.Time  `json:"createTime,omitempty"`
	CreateTimeNEQ   *time.Time  `json:"createTimeNEQ,omitempty"`
	CreateTimeIn    []time.Time `json:"createTimeIn,omitempty"`
	CreateTimeNotIn []time.Time `json:"createTimeNotIn,omitempty"`
	CreateTimeGT    *time.Time  `json:"createTimeGT,omitempty"`
	CreateTimeGTE   *time.Time  `json:"createTimeGTE,omitempty"`
	CreateTimeLT    *time.Time  `json:"createTimeLT,omitempty"`
	CreateTimeLTE   *time.Time  `json:"createTimeLTE,omitempty"`

	// "update_time" field predicates.
	UpdateTime      *time.Time  `json:"updateTime,omitempty"`
	UpdateTimeNEQ   *time.Time  `json:"updateTimeNEQ,omitempty"`
	UpdateTimeIn    []time.Time `json:"updateTimeIn,omitempty"`
	UpdateTimeNotIn []time.Time `json:"updateTimeNotIn,omitempty"`
	UpdateTimeGT    *time.Time  `json:"updateTimeGT,omitempty"`
	UpdateTimeGTE   *time.Time  `json:"updateTimeGTE,omitempty"`
	UpdateTimeLT    *time.Time  `json:"updateTimeLT,omitempty"`
	UpdateTimeLTE   *time.Time  `json:"updateTimeLTE,omitempty"`

	// "title" field predicates.
	Title             *string  `json:"title,omitempty"`
	TitleNEQ          *string  `json:"titleNEQ,omitempty"`
	TitleIn           []string `json:"titleIn,omitempty"`
	TitleNotIn        []string `json:"titleNotIn,omitempty"`
	TitleGT           *string  `json:"titleGT,omitempty"`
	TitleGTE          *string  `json:"titleGTE,omitempty"`
	TitleLT           *string  `json:"titleLT,omitempty"`
	TitleLTE          *string  `json:"titleLTE,omitempty"`
	TitleContains     *string  `json:"titleContains,omitempty"`
	TitleHasPrefix    *string  `json:"titleHasPrefix,omitempty"`
	TitleHasSuffix    *string  `json:"titleHasSuffix,omitempty"`
	TitleEqualFold    *string  `json:"titleEqualFold,omitempty"`
	TitleContainsFold *string  `json:"titleContainsFold,omitempty"`

	// "content" field predicates.
	Content             *string  `json:"content,omitempty"`
	ContentNEQ          *string  `json:"contentNEQ,omitempty"`
	ContentIn           []string `json:"contentIn,omitempty"`
	ContentNotIn        []string `json:"contentNotIn,omitempty"`
	ContentGT           *string  `json:"contentGT,omitempty"`
	ContentGTE          *string  `json:"contentGTE,omitempty"`
	ContentLT           *string  `json:"contentLT,omitempty"`
	ContentLTE          *string  `json:"contentLTE,omitempty"`
	ContentContains     *string  `json:"contentContains,omitempty"`
	ContentHasPrefix    *string  `json:"contentHasPrefix,omitempty"`
	ContentHasSuffix    *string  `json:"contentHasSuffix,omitempty"`
	ContentEqualFold    *string  `json:"contentEqualFold,omitempty"`
	ContentContainsFold *string  `json:"contentContainsFold,omitempty"`

	// "type" field predicates.
	Type      *post.Type  `json:"type,omitempty"`
	TypeNEQ   *post.Type  `json:"typeNEQ,omitempty"`
	TypeIn    []post.Type `json:"typeIn,omitempty"`
	TypeNotIn []post.Type `json:"typeNotIn,omitempty"`

	// "owner" edge predicates.
	HasOwner     *bool             `json:"hasOwner,omitempty"`
	HasOwnerWith []*UserWhereInput `json:"hasOwnerWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *PostWhereInput) AddPredicates(predicates ...predicate.Post) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the PostWhereInput filter on the PostQuery builder.
func (i *PostWhereInput) Filter(q *PostQuery) (*PostQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering posts.
// An error is returned if the input is empty or invalid.
func (i *PostWhereInput) P() (predicate.Post, error) {
	var predicates []predicate.Post
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, post.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Post, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, post.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Post, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, post.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, post.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, post.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, post.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, post.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, post.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, post.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, post.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, post.IDLTE(*i.IDLTE))
	}
	if i.CreateTime != nil {
		predicates = append(predicates, post.CreateTimeEQ(*i.CreateTime))
	}
	if i.CreateTimeNEQ != nil {
		predicates = append(predicates, post.CreateTimeNEQ(*i.CreateTimeNEQ))
	}
	if len(i.CreateTimeIn) > 0 {
		predicates = append(predicates, post.CreateTimeIn(i.CreateTimeIn...))
	}
	if len(i.CreateTimeNotIn) > 0 {
		predicates = append(predicates, post.CreateTimeNotIn(i.CreateTimeNotIn...))
	}
	if i.CreateTimeGT != nil {
		predicates = append(predicates, post.CreateTimeGT(*i.CreateTimeGT))
	}
	if i.CreateTimeGTE != nil {
		predicates = append(predicates, post.CreateTimeGTE(*i.CreateTimeGTE))
	}
	if i.CreateTimeLT != nil {
		predicates = append(predicates, post.CreateTimeLT(*i.CreateTimeLT))
	}
	if i.CreateTimeLTE != nil {
		predicates = append(predicates, post.CreateTimeLTE(*i.CreateTimeLTE))
	}
	if i.UpdateTime != nil {
		predicates = append(predicates, post.UpdateTimeEQ(*i.UpdateTime))
	}
	if i.UpdateTimeNEQ != nil {
		predicates = append(predicates, post.UpdateTimeNEQ(*i.UpdateTimeNEQ))
	}
	if len(i.UpdateTimeIn) > 0 {
		predicates = append(predicates, post.UpdateTimeIn(i.UpdateTimeIn...))
	}
	if len(i.UpdateTimeNotIn) > 0 {
		predicates = append(predicates, post.UpdateTimeNotIn(i.UpdateTimeNotIn...))
	}
	if i.UpdateTimeGT != nil {
		predicates = append(predicates, post.UpdateTimeGT(*i.UpdateTimeGT))
	}
	if i.UpdateTimeGTE != nil {
		predicates = append(predicates, post.UpdateTimeGTE(*i.UpdateTimeGTE))
	}
	if i.UpdateTimeLT != nil {
		predicates = append(predicates, post.UpdateTimeLT(*i.UpdateTimeLT))
	}
	if i.UpdateTimeLTE != nil {
		predicates = append(predicates, post.UpdateTimeLTE(*i.UpdateTimeLTE))
	}
	if i.Title != nil {
		predicates = append(predicates, post.TitleEQ(*i.Title))
	}
	if i.TitleNEQ != nil {
		predicates = append(predicates, post.TitleNEQ(*i.TitleNEQ))
	}
	if len(i.TitleIn) > 0 {
		predicates = append(predicates, post.TitleIn(i.TitleIn...))
	}
	if len(i.TitleNotIn) > 0 {
		predicates = append(predicates, post.TitleNotIn(i.TitleNotIn...))
	}
	if i.TitleGT != nil {
		predicates = append(predicates, post.TitleGT(*i.TitleGT))
	}
	if i.TitleGTE != nil {
		predicates = append(predicates, post.TitleGTE(*i.TitleGTE))
	}
	if i.TitleLT != nil {
		predicates = append(predicates, post.TitleLT(*i.TitleLT))
	}
	if i.TitleLTE != nil {
		predicates = append(predicates, post.TitleLTE(*i.TitleLTE))
	}
	if i.TitleContains != nil {
		predicates = append(predicates, post.TitleContains(*i.TitleContains))
	}
	if i.TitleHasPrefix != nil {
		predicates = append(predicates, post.TitleHasPrefix(*i.TitleHasPrefix))
	}
	if i.TitleHasSuffix != nil {
		predicates = append(predicates, post.TitleHasSuffix(*i.TitleHasSuffix))
	}
	if i.TitleEqualFold != nil {
		predicates = append(predicates, post.TitleEqualFold(*i.TitleEqualFold))
	}
	if i.TitleContainsFold != nil {
		predicates = append(predicates, post.TitleContainsFold(*i.TitleContainsFold))
	}
	if i.Content != nil {
		predicates = append(predicates, post.ContentEQ(*i.Content))
	}
	if i.ContentNEQ != nil {
		predicates = append(predicates, post.ContentNEQ(*i.ContentNEQ))
	}
	if len(i.ContentIn) > 0 {
		predicates = append(predicates, post.ContentIn(i.ContentIn...))
	}
	if len(i.ContentNotIn) > 0 {
		predicates = append(predicates, post.ContentNotIn(i.ContentNotIn...))
	}
	if i.ContentGT != nil {
		predicates = append(predicates, post.ContentGT(*i.ContentGT))
	}
	if i.ContentGTE != nil {
		predicates = append(predicates, post.ContentGTE(*i.ContentGTE))
	}
	if i.ContentLT != nil {
		predicates = append(predicates, post.ContentLT(*i.ContentLT))
	}
	if i.ContentLTE != nil {
		predicates = append(predicates, post.ContentLTE(*i.ContentLTE))
	}
	if i.ContentContains != nil {
		predicates = append(predicates, post.ContentContains(*i.ContentContains))
	}
	if i.ContentHasPrefix != nil {
		predicates = append(predicates, post.ContentHasPrefix(*i.ContentHasPrefix))
	}
	if i.ContentHasSuffix != nil {
		predicates = append(predicates, post.ContentHasSuffix(*i.ContentHasSuffix))
	}
	if i.ContentEqualFold != nil {
		predicates = append(predicates, post.ContentEqualFold(*i.ContentEqualFold))
	}
	if i.ContentContainsFold != nil {
		predicates = append(predicates, post.ContentContainsFold(*i.ContentContainsFold))
	}
	if i.Type != nil {
		predicates = append(predicates, post.TypeEQ(*i.Type))
	}
	if i.TypeNEQ != nil {
		predicates = append(predicates, post.TypeNEQ(*i.TypeNEQ))
	}
	if len(i.TypeIn) > 0 {
		predicates = append(predicates, post.TypeIn(i.TypeIn...))
	}
	if len(i.TypeNotIn) > 0 {
		predicates = append(predicates, post.TypeNotIn(i.TypeNotIn...))
	}

	if i.HasOwner != nil {
		p := post.HasOwner()
		if !*i.HasOwner {
			p = post.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOwnerWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasOwnerWith))
		for _, w := range i.HasOwnerWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, post.HasOwnerWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("empty predicate PostWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return post.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "posts" edge predicates.
	HasPosts     *bool             `json:"hasPosts,omitempty"`
	HasPostsWith []*PostWhereInput `json:"hasPostsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasPosts != nil {
		p := user.HasPosts()
		if !*i.HasPosts {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasPostsWith) > 0 {
		with := make([]predicate.Post, 0, len(i.HasPostsWith))
		for _, w := range i.HasPostsWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasPostsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("empty predicate UserWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}