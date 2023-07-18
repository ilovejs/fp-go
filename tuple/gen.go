// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2023-07-18 15:21:49.4411816 +0200 CEST m=+0.105633301
package tuple


import (
	M "github.com/ibm/fp-go/monoid"
	O "github.com/ibm/fp-go/ord"
)

// Tuple1 is a struct that carries 1 independently typed values
type Tuple1[T1 any] struct {
  F1 T1
}

// Tuple2 is a struct that carries 2 independently typed values
type Tuple2[T1, T2 any] struct {
  F1 T1
  F2 T2
}

// Tuple3 is a struct that carries 3 independently typed values
type Tuple3[T1, T2, T3 any] struct {
  F1 T1
  F2 T2
  F3 T3
}

// Tuple4 is a struct that carries 4 independently typed values
type Tuple4[T1, T2, T3, T4 any] struct {
  F1 T1
  F2 T2
  F3 T3
  F4 T4
}

// Tuple5 is a struct that carries 5 independently typed values
type Tuple5[T1, T2, T3, T4, T5 any] struct {
  F1 T1
  F2 T2
  F3 T3
  F4 T4
  F5 T5
}

// Tuple6 is a struct that carries 6 independently typed values
type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
  F1 T1
  F2 T2
  F3 T3
  F4 T4
  F5 T5
  F6 T6
}

// Tuple7 is a struct that carries 7 independently typed values
type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
  F1 T1
  F2 T2
  F3 T3
  F4 T4
  F5 T5
  F6 T6
  F7 T7
}

// Tuple8 is a struct that carries 8 independently typed values
type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
  F1 T1
  F2 T2
  F3 T3
  F4 T4
  F5 T5
  F6 T6
  F7 T7
  F8 T8
}

// Tuple9 is a struct that carries 9 independently typed values
type Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
  F1 T1
  F2 T2
  F3 T3
  F4 T4
  F5 T5
  F6 T6
  F7 T7
  F8 T8
  F9 T9
}

// Tuple10 is a struct that carries 10 independently typed values
type Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
  F1 T1
  F2 T2
  F3 T3
  F4 T4
  F5 T5
  F6 T6
  F7 T7
  F8 T8
  F9 T9
  F10 T10
}

// MakeTuple1 is a function that converts its 1 parameters into a [Tuple1]
func MakeTuple1[T1 any](t1 T1) Tuple1[T1] {
  return Tuple1[T1]{t1}
}

// Tupled1 converts a function with 1 parameters returning into a function taking a Tuple1
// The inverse function is [Untupled1]
func Tupled1[F ~func(T1) R, T1, R any](f F) func(Tuple1[T1]) R {
  return func(t Tuple1[T1]) R {
    return f(t.F1)
  }
}

// Untupled1 converts a function with a [Tuple1] parameter into a function with 1 parameters
// The inverse function is [Tupled1]
func Untupled1[F ~func(Tuple1[T1]) R, T1, R any](f F) func(T1) R {
  return func(t1 T1) R {
    return f(MakeTuple1(t1))
  }
}

// Monoid1 creates a [Monoid] for a [Tuple1] based on 1 monoids for the contained types
func Monoid1[T1 any](m1 M.Monoid[T1]) M.Monoid[Tuple1[T1]] {
  return M.MakeMonoid(func(l, r Tuple1[T1]) Tuple1[T1]{
    return MakeTuple1(m1.Concat(l.F1, r.F1))
  }, MakeTuple1(m1.Empty()))
}

// Ord1 creates n [Ord] for a [Tuple1] based on 1 [Ord]s for the contained types
func Ord1[T1 any](o1 O.Ord[T1]) O.Ord[Tuple1[T1]] {
  return O.MakeOrd(func(l, r Tuple1[T1]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    return 0
  }, func(l, r Tuple1[T1]) bool {
    return o1.Equals(l.F1, r.F1)
  })
}

// Map1 maps each value of a [Tuple1] via a mapping function
func Map1[F1 ~func(T1) R1, T1, R1 any](f1 F1) func(Tuple1[T1]) Tuple1[R1] {
 return func(t Tuple1[T1]) Tuple1[R1] {
    return MakeTuple1(
      f1(t.F1),
    )
 }
}

// Replicate1 creates a [Tuple1] with all fields set to the input value `t`
func Replicate1[T any](t T) Tuple1[T] {
  return MakeTuple1(t)
}

// MakeTuple2 is a function that converts its 2 parameters into a [Tuple2]
func MakeTuple2[T1, T2 any](t1 T1, t2 T2) Tuple2[T1, T2] {
  return Tuple2[T1, T2]{t1, t2}
}

// Tupled2 converts a function with 2 parameters returning into a function taking a Tuple2
// The inverse function is [Untupled2]
func Tupled2[F ~func(T1, T2) R, T1, T2, R any](f F) func(Tuple2[T1, T2]) R {
  return func(t Tuple2[T1, T2]) R {
    return f(t.F1, t.F2)
  }
}

// Untupled2 converts a function with a [Tuple2] parameter into a function with 2 parameters
// The inverse function is [Tupled2]
func Untupled2[F ~func(Tuple2[T1, T2]) R, T1, T2, R any](f F) func(T1, T2) R {
  return func(t1 T1, t2 T2) R {
    return f(MakeTuple2(t1, t2))
  }
}

// Monoid2 creates a [Monoid] for a [Tuple2] based on 2 monoids for the contained types
func Monoid2[T1, T2 any](m1 M.Monoid[T1], m2 M.Monoid[T2]) M.Monoid[Tuple2[T1, T2]] {
  return M.MakeMonoid(func(l, r Tuple2[T1, T2]) Tuple2[T1, T2]{
    return MakeTuple2(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2))
  }, MakeTuple2(m1.Empty(), m2.Empty()))
}

// Ord2 creates n [Ord] for a [Tuple2] based on 2 [Ord]s for the contained types
func Ord2[T1, T2 any](o1 O.Ord[T1], o2 O.Ord[T2]) O.Ord[Tuple2[T1, T2]] {
  return O.MakeOrd(func(l, r Tuple2[T1, T2]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    return 0
  }, func(l, r Tuple2[T1, T2]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2)
  })
}

// Map2 maps each value of a [Tuple2] via a mapping function
func Map2[F1 ~func(T1) R1, F2 ~func(T2) R2, T1, R1, T2, R2 any](f1 F1, f2 F2) func(Tuple2[T1, T2]) Tuple2[R1, R2] {
 return func(t Tuple2[T1, T2]) Tuple2[R1, R2] {
    return MakeTuple2(
      f1(t.F1),
      f2(t.F2),
    )
 }
}

// Replicate2 creates a [Tuple2] with all fields set to the input value `t`
func Replicate2[T any](t T) Tuple2[T, T] {
  return MakeTuple2(t, t)
}

// MakeTuple3 is a function that converts its 3 parameters into a [Tuple3]
func MakeTuple3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3) Tuple3[T1, T2, T3] {
  return Tuple3[T1, T2, T3]{t1, t2, t3}
}

// Tupled3 converts a function with 3 parameters returning into a function taking a Tuple3
// The inverse function is [Untupled3]
func Tupled3[F ~func(T1, T2, T3) R, T1, T2, T3, R any](f F) func(Tuple3[T1, T2, T3]) R {
  return func(t Tuple3[T1, T2, T3]) R {
    return f(t.F1, t.F2, t.F3)
  }
}

// Untupled3 converts a function with a [Tuple3] parameter into a function with 3 parameters
// The inverse function is [Tupled3]
func Untupled3[F ~func(Tuple3[T1, T2, T3]) R, T1, T2, T3, R any](f F) func(T1, T2, T3) R {
  return func(t1 T1, t2 T2, t3 T3) R {
    return f(MakeTuple3(t1, t2, t3))
  }
}

// Monoid3 creates a [Monoid] for a [Tuple3] based on 3 monoids for the contained types
func Monoid3[T1, T2, T3 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3]) M.Monoid[Tuple3[T1, T2, T3]] {
  return M.MakeMonoid(func(l, r Tuple3[T1, T2, T3]) Tuple3[T1, T2, T3]{
    return MakeTuple3(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3))
  }, MakeTuple3(m1.Empty(), m2.Empty(), m3.Empty()))
}

// Ord3 creates n [Ord] for a [Tuple3] based on 3 [Ord]s for the contained types
func Ord3[T1, T2, T3 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3]) O.Ord[Tuple3[T1, T2, T3]] {
  return O.MakeOrd(func(l, r Tuple3[T1, T2, T3]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    return 0
  }, func(l, r Tuple3[T1, T2, T3]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3)
  })
}

// Map3 maps each value of a [Tuple3] via a mapping function
func Map3[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, T1, R1, T2, R2, T3, R3 any](f1 F1, f2 F2, f3 F3) func(Tuple3[T1, T2, T3]) Tuple3[R1, R2, R3] {
 return func(t Tuple3[T1, T2, T3]) Tuple3[R1, R2, R3] {
    return MakeTuple3(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
    )
 }
}

// Replicate3 creates a [Tuple3] with all fields set to the input value `t`
func Replicate3[T any](t T) Tuple3[T, T, T] {
  return MakeTuple3(t, t, t)
}

// MakeTuple4 is a function that converts its 4 parameters into a [Tuple4]
func MakeTuple4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4) Tuple4[T1, T2, T3, T4] {
  return Tuple4[T1, T2, T3, T4]{t1, t2, t3, t4}
}

// Tupled4 converts a function with 4 parameters returning into a function taking a Tuple4
// The inverse function is [Untupled4]
func Tupled4[F ~func(T1, T2, T3, T4) R, T1, T2, T3, T4, R any](f F) func(Tuple4[T1, T2, T3, T4]) R {
  return func(t Tuple4[T1, T2, T3, T4]) R {
    return f(t.F1, t.F2, t.F3, t.F4)
  }
}

// Untupled4 converts a function with a [Tuple4] parameter into a function with 4 parameters
// The inverse function is [Tupled4]
func Untupled4[F ~func(Tuple4[T1, T2, T3, T4]) R, T1, T2, T3, T4, R any](f F) func(T1, T2, T3, T4) R {
  return func(t1 T1, t2 T2, t3 T3, t4 T4) R {
    return f(MakeTuple4(t1, t2, t3, t4))
  }
}

// Monoid4 creates a [Monoid] for a [Tuple4] based on 4 monoids for the contained types
func Monoid4[T1, T2, T3, T4 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3], m4 M.Monoid[T4]) M.Monoid[Tuple4[T1, T2, T3, T4]] {
  return M.MakeMonoid(func(l, r Tuple4[T1, T2, T3, T4]) Tuple4[T1, T2, T3, T4]{
    return MakeTuple4(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3), m4.Concat(l.F4, r.F4))
  }, MakeTuple4(m1.Empty(), m2.Empty(), m3.Empty(), m4.Empty()))
}

// Ord4 creates n [Ord] for a [Tuple4] based on 4 [Ord]s for the contained types
func Ord4[T1, T2, T3, T4 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3], o4 O.Ord[T4]) O.Ord[Tuple4[T1, T2, T3, T4]] {
  return O.MakeOrd(func(l, r Tuple4[T1, T2, T3, T4]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    if c:= o4.Compare(l.F4, r.F4); c != 0 {return c}
    return 0
  }, func(l, r Tuple4[T1, T2, T3, T4]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3) && o4.Equals(l.F4, r.F4)
  })
}

// Map4 maps each value of a [Tuple4] via a mapping function
func Map4[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, F4 ~func(T4) R4, T1, R1, T2, R2, T3, R3, T4, R4 any](f1 F1, f2 F2, f3 F3, f4 F4) func(Tuple4[T1, T2, T3, T4]) Tuple4[R1, R2, R3, R4] {
 return func(t Tuple4[T1, T2, T3, T4]) Tuple4[R1, R2, R3, R4] {
    return MakeTuple4(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
      f4(t.F4),
    )
 }
}

// Replicate4 creates a [Tuple4] with all fields set to the input value `t`
func Replicate4[T any](t T) Tuple4[T, T, T, T] {
  return MakeTuple4(t, t, t, t)
}

// MakeTuple5 is a function that converts its 5 parameters into a [Tuple5]
func MakeTuple5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) Tuple5[T1, T2, T3, T4, T5] {
  return Tuple5[T1, T2, T3, T4, T5]{t1, t2, t3, t4, t5}
}

// Tupled5 converts a function with 5 parameters returning into a function taking a Tuple5
// The inverse function is [Untupled5]
func Tupled5[F ~func(T1, T2, T3, T4, T5) R, T1, T2, T3, T4, T5, R any](f F) func(Tuple5[T1, T2, T3, T4, T5]) R {
  return func(t Tuple5[T1, T2, T3, T4, T5]) R {
    return f(t.F1, t.F2, t.F3, t.F4, t.F5)
  }
}

// Untupled5 converts a function with a [Tuple5] parameter into a function with 5 parameters
// The inverse function is [Tupled5]
func Untupled5[F ~func(Tuple5[T1, T2, T3, T4, T5]) R, T1, T2, T3, T4, T5, R any](f F) func(T1, T2, T3, T4, T5) R {
  return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5) R {
    return f(MakeTuple5(t1, t2, t3, t4, t5))
  }
}

// Monoid5 creates a [Monoid] for a [Tuple5] based on 5 monoids for the contained types
func Monoid5[T1, T2, T3, T4, T5 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3], m4 M.Monoid[T4], m5 M.Monoid[T5]) M.Monoid[Tuple5[T1, T2, T3, T4, T5]] {
  return M.MakeMonoid(func(l, r Tuple5[T1, T2, T3, T4, T5]) Tuple5[T1, T2, T3, T4, T5]{
    return MakeTuple5(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3), m4.Concat(l.F4, r.F4), m5.Concat(l.F5, r.F5))
  }, MakeTuple5(m1.Empty(), m2.Empty(), m3.Empty(), m4.Empty(), m5.Empty()))
}

// Ord5 creates n [Ord] for a [Tuple5] based on 5 [Ord]s for the contained types
func Ord5[T1, T2, T3, T4, T5 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3], o4 O.Ord[T4], o5 O.Ord[T5]) O.Ord[Tuple5[T1, T2, T3, T4, T5]] {
  return O.MakeOrd(func(l, r Tuple5[T1, T2, T3, T4, T5]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    if c:= o4.Compare(l.F4, r.F4); c != 0 {return c}
    if c:= o5.Compare(l.F5, r.F5); c != 0 {return c}
    return 0
  }, func(l, r Tuple5[T1, T2, T3, T4, T5]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3) && o4.Equals(l.F4, r.F4) && o5.Equals(l.F5, r.F5)
  })
}

// Map5 maps each value of a [Tuple5] via a mapping function
func Map5[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, F4 ~func(T4) R4, F5 ~func(T5) R5, T1, R1, T2, R2, T3, R3, T4, R4, T5, R5 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5) func(Tuple5[T1, T2, T3, T4, T5]) Tuple5[R1, R2, R3, R4, R5] {
 return func(t Tuple5[T1, T2, T3, T4, T5]) Tuple5[R1, R2, R3, R4, R5] {
    return MakeTuple5(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
      f4(t.F4),
      f5(t.F5),
    )
 }
}

// Replicate5 creates a [Tuple5] with all fields set to the input value `t`
func Replicate5[T any](t T) Tuple5[T, T, T, T, T] {
  return MakeTuple5(t, t, t, t, t)
}

// MakeTuple6 is a function that converts its 6 parameters into a [Tuple6]
func MakeTuple6[T1, T2, T3, T4, T5, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) Tuple6[T1, T2, T3, T4, T5, T6] {
  return Tuple6[T1, T2, T3, T4, T5, T6]{t1, t2, t3, t4, t5, t6}
}

// Tupled6 converts a function with 6 parameters returning into a function taking a Tuple6
// The inverse function is [Untupled6]
func Tupled6[F ~func(T1, T2, T3, T4, T5, T6) R, T1, T2, T3, T4, T5, T6, R any](f F) func(Tuple6[T1, T2, T3, T4, T5, T6]) R {
  return func(t Tuple6[T1, T2, T3, T4, T5, T6]) R {
    return f(t.F1, t.F2, t.F3, t.F4, t.F5, t.F6)
  }
}

// Untupled6 converts a function with a [Tuple6] parameter into a function with 6 parameters
// The inverse function is [Tupled6]
func Untupled6[F ~func(Tuple6[T1, T2, T3, T4, T5, T6]) R, T1, T2, T3, T4, T5, T6, R any](f F) func(T1, T2, T3, T4, T5, T6) R {
  return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) R {
    return f(MakeTuple6(t1, t2, t3, t4, t5, t6))
  }
}

// Monoid6 creates a [Monoid] for a [Tuple6] based on 6 monoids for the contained types
func Monoid6[T1, T2, T3, T4, T5, T6 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3], m4 M.Monoid[T4], m5 M.Monoid[T5], m6 M.Monoid[T6]) M.Monoid[Tuple6[T1, T2, T3, T4, T5, T6]] {
  return M.MakeMonoid(func(l, r Tuple6[T1, T2, T3, T4, T5, T6]) Tuple6[T1, T2, T3, T4, T5, T6]{
    return MakeTuple6(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3), m4.Concat(l.F4, r.F4), m5.Concat(l.F5, r.F5), m6.Concat(l.F6, r.F6))
  }, MakeTuple6(m1.Empty(), m2.Empty(), m3.Empty(), m4.Empty(), m5.Empty(), m6.Empty()))
}

// Ord6 creates n [Ord] for a [Tuple6] based on 6 [Ord]s for the contained types
func Ord6[T1, T2, T3, T4, T5, T6 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3], o4 O.Ord[T4], o5 O.Ord[T5], o6 O.Ord[T6]) O.Ord[Tuple6[T1, T2, T3, T4, T5, T6]] {
  return O.MakeOrd(func(l, r Tuple6[T1, T2, T3, T4, T5, T6]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    if c:= o4.Compare(l.F4, r.F4); c != 0 {return c}
    if c:= o5.Compare(l.F5, r.F5); c != 0 {return c}
    if c:= o6.Compare(l.F6, r.F6); c != 0 {return c}
    return 0
  }, func(l, r Tuple6[T1, T2, T3, T4, T5, T6]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3) && o4.Equals(l.F4, r.F4) && o5.Equals(l.F5, r.F5) && o6.Equals(l.F6, r.F6)
  })
}

// Map6 maps each value of a [Tuple6] via a mapping function
func Map6[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, F4 ~func(T4) R4, F5 ~func(T5) R5, F6 ~func(T6) R6, T1, R1, T2, R2, T3, R3, T4, R4, T5, R5, T6, R6 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6) func(Tuple6[T1, T2, T3, T4, T5, T6]) Tuple6[R1, R2, R3, R4, R5, R6] {
 return func(t Tuple6[T1, T2, T3, T4, T5, T6]) Tuple6[R1, R2, R3, R4, R5, R6] {
    return MakeTuple6(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
      f4(t.F4),
      f5(t.F5),
      f6(t.F6),
    )
 }
}

// Replicate6 creates a [Tuple6] with all fields set to the input value `t`
func Replicate6[T any](t T) Tuple6[T, T, T, T, T, T] {
  return MakeTuple6(t, t, t, t, t, t)
}

// MakeTuple7 is a function that converts its 7 parameters into a [Tuple7]
func MakeTuple7[T1, T2, T3, T4, T5, T6, T7 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) Tuple7[T1, T2, T3, T4, T5, T6, T7] {
  return Tuple7[T1, T2, T3, T4, T5, T6, T7]{t1, t2, t3, t4, t5, t6, t7}
}

// Tupled7 converts a function with 7 parameters returning into a function taking a Tuple7
// The inverse function is [Untupled7]
func Tupled7[F ~func(T1, T2, T3, T4, T5, T6, T7) R, T1, T2, T3, T4, T5, T6, T7, R any](f F) func(Tuple7[T1, T2, T3, T4, T5, T6, T7]) R {
  return func(t Tuple7[T1, T2, T3, T4, T5, T6, T7]) R {
    return f(t.F1, t.F2, t.F3, t.F4, t.F5, t.F6, t.F7)
  }
}

// Untupled7 converts a function with a [Tuple7] parameter into a function with 7 parameters
// The inverse function is [Tupled7]
func Untupled7[F ~func(Tuple7[T1, T2, T3, T4, T5, T6, T7]) R, T1, T2, T3, T4, T5, T6, T7, R any](f F) func(T1, T2, T3, T4, T5, T6, T7) R {
  return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7) R {
    return f(MakeTuple7(t1, t2, t3, t4, t5, t6, t7))
  }
}

// Monoid7 creates a [Monoid] for a [Tuple7] based on 7 monoids for the contained types
func Monoid7[T1, T2, T3, T4, T5, T6, T7 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3], m4 M.Monoid[T4], m5 M.Monoid[T5], m6 M.Monoid[T6], m7 M.Monoid[T7]) M.Monoid[Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
  return M.MakeMonoid(func(l, r Tuple7[T1, T2, T3, T4, T5, T6, T7]) Tuple7[T1, T2, T3, T4, T5, T6, T7]{
    return MakeTuple7(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3), m4.Concat(l.F4, r.F4), m5.Concat(l.F5, r.F5), m6.Concat(l.F6, r.F6), m7.Concat(l.F7, r.F7))
  }, MakeTuple7(m1.Empty(), m2.Empty(), m3.Empty(), m4.Empty(), m5.Empty(), m6.Empty(), m7.Empty()))
}

// Ord7 creates n [Ord] for a [Tuple7] based on 7 [Ord]s for the contained types
func Ord7[T1, T2, T3, T4, T5, T6, T7 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3], o4 O.Ord[T4], o5 O.Ord[T5], o6 O.Ord[T6], o7 O.Ord[T7]) O.Ord[Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
  return O.MakeOrd(func(l, r Tuple7[T1, T2, T3, T4, T5, T6, T7]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    if c:= o4.Compare(l.F4, r.F4); c != 0 {return c}
    if c:= o5.Compare(l.F5, r.F5); c != 0 {return c}
    if c:= o6.Compare(l.F6, r.F6); c != 0 {return c}
    if c:= o7.Compare(l.F7, r.F7); c != 0 {return c}
    return 0
  }, func(l, r Tuple7[T1, T2, T3, T4, T5, T6, T7]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3) && o4.Equals(l.F4, r.F4) && o5.Equals(l.F5, r.F5) && o6.Equals(l.F6, r.F6) && o7.Equals(l.F7, r.F7)
  })
}

// Map7 maps each value of a [Tuple7] via a mapping function
func Map7[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, F4 ~func(T4) R4, F5 ~func(T5) R5, F6 ~func(T6) R6, F7 ~func(T7) R7, T1, R1, T2, R2, T3, R3, T4, R4, T5, R5, T6, R6, T7, R7 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7) func(Tuple7[T1, T2, T3, T4, T5, T6, T7]) Tuple7[R1, R2, R3, R4, R5, R6, R7] {
 return func(t Tuple7[T1, T2, T3, T4, T5, T6, T7]) Tuple7[R1, R2, R3, R4, R5, R6, R7] {
    return MakeTuple7(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
      f4(t.F4),
      f5(t.F5),
      f6(t.F6),
      f7(t.F7),
    )
 }
}

// Replicate7 creates a [Tuple7] with all fields set to the input value `t`
func Replicate7[T any](t T) Tuple7[T, T, T, T, T, T, T] {
  return MakeTuple7(t, t, t, t, t, t, t)
}

// MakeTuple8 is a function that converts its 8 parameters into a [Tuple8]
func MakeTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
  return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{t1, t2, t3, t4, t5, t6, t7, t8}
}

// Tupled8 converts a function with 8 parameters returning into a function taking a Tuple8
// The inverse function is [Untupled8]
func Tupled8[F ~func(T1, T2, T3, T4, T5, T6, T7, T8) R, T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) R {
  return func(t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) R {
    return f(t.F1, t.F2, t.F3, t.F4, t.F5, t.F6, t.F7, t.F8)
  }
}

// Untupled8 converts a function with a [Tuple8] parameter into a function with 8 parameters
// The inverse function is [Tupled8]
func Untupled8[F ~func(Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) R, T1, T2, T3, T4, T5, T6, T7, T8, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8) R {
  return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8) R {
    return f(MakeTuple8(t1, t2, t3, t4, t5, t6, t7, t8))
  }
}

// Monoid8 creates a [Monoid] for a [Tuple8] based on 8 monoids for the contained types
func Monoid8[T1, T2, T3, T4, T5, T6, T7, T8 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3], m4 M.Monoid[T4], m5 M.Monoid[T5], m6 M.Monoid[T6], m7 M.Monoid[T7], m8 M.Monoid[T8]) M.Monoid[Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
  return M.MakeMonoid(func(l, r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
    return MakeTuple8(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3), m4.Concat(l.F4, r.F4), m5.Concat(l.F5, r.F5), m6.Concat(l.F6, r.F6), m7.Concat(l.F7, r.F7), m8.Concat(l.F8, r.F8))
  }, MakeTuple8(m1.Empty(), m2.Empty(), m3.Empty(), m4.Empty(), m5.Empty(), m6.Empty(), m7.Empty(), m8.Empty()))
}

// Ord8 creates n [Ord] for a [Tuple8] based on 8 [Ord]s for the contained types
func Ord8[T1, T2, T3, T4, T5, T6, T7, T8 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3], o4 O.Ord[T4], o5 O.Ord[T5], o6 O.Ord[T6], o7 O.Ord[T7], o8 O.Ord[T8]) O.Ord[Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
  return O.MakeOrd(func(l, r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    if c:= o4.Compare(l.F4, r.F4); c != 0 {return c}
    if c:= o5.Compare(l.F5, r.F5); c != 0 {return c}
    if c:= o6.Compare(l.F6, r.F6); c != 0 {return c}
    if c:= o7.Compare(l.F7, r.F7); c != 0 {return c}
    if c:= o8.Compare(l.F8, r.F8); c != 0 {return c}
    return 0
  }, func(l, r Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3) && o4.Equals(l.F4, r.F4) && o5.Equals(l.F5, r.F5) && o6.Equals(l.F6, r.F6) && o7.Equals(l.F7, r.F7) && o8.Equals(l.F8, r.F8)
  })
}

// Map8 maps each value of a [Tuple8] via a mapping function
func Map8[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, F4 ~func(T4) R4, F5 ~func(T5) R5, F6 ~func(T6) R6, F7 ~func(T7) R7, F8 ~func(T8) R8, T1, R1, T2, R2, T3, R3, T4, R4, T5, R5, T6, R6, T7, R7, T8, R8 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8) func(Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Tuple8[R1, R2, R3, R4, R5, R6, R7, R8] {
 return func(t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Tuple8[R1, R2, R3, R4, R5, R6, R7, R8] {
    return MakeTuple8(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
      f4(t.F4),
      f5(t.F5),
      f6(t.F6),
      f7(t.F7),
      f8(t.F8),
    )
 }
}

// Replicate8 creates a [Tuple8] with all fields set to the input value `t`
func Replicate8[T any](t T) Tuple8[T, T, T, T, T, T, T, T] {
  return MakeTuple8(t, t, t, t, t, t, t, t)
}

// MakeTuple9 is a function that converts its 9 parameters into a [Tuple9]
func MakeTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
  return Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{t1, t2, t3, t4, t5, t6, t7, t8, t9}
}

// Tupled9 converts a function with 9 parameters returning into a function taking a Tuple9
// The inverse function is [Untupled9]
func Tupled9[F ~func(T1, T2, T3, T4, T5, T6, T7, T8, T9) R, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) R {
  return func(t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) R {
    return f(t.F1, t.F2, t.F3, t.F4, t.F5, t.F6, t.F7, t.F8, t.F9)
  }
}

// Untupled9 converts a function with a [Tuple9] parameter into a function with 9 parameters
// The inverse function is [Tupled9]
func Untupled9[F ~func(Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) R, T1, T2, T3, T4, T5, T6, T7, T8, T9, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8, T9) R {
  return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9) R {
    return f(MakeTuple9(t1, t2, t3, t4, t5, t6, t7, t8, t9))
  }
}

// Monoid9 creates a [Monoid] for a [Tuple9] based on 9 monoids for the contained types
func Monoid9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3], m4 M.Monoid[T4], m5 M.Monoid[T5], m6 M.Monoid[T6], m7 M.Monoid[T7], m8 M.Monoid[T8], m9 M.Monoid[T9]) M.Monoid[Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
  return M.MakeMonoid(func(l, r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
    return MakeTuple9(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3), m4.Concat(l.F4, r.F4), m5.Concat(l.F5, r.F5), m6.Concat(l.F6, r.F6), m7.Concat(l.F7, r.F7), m8.Concat(l.F8, r.F8), m9.Concat(l.F9, r.F9))
  }, MakeTuple9(m1.Empty(), m2.Empty(), m3.Empty(), m4.Empty(), m5.Empty(), m6.Empty(), m7.Empty(), m8.Empty(), m9.Empty()))
}

// Ord9 creates n [Ord] for a [Tuple9] based on 9 [Ord]s for the contained types
func Ord9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3], o4 O.Ord[T4], o5 O.Ord[T5], o6 O.Ord[T6], o7 O.Ord[T7], o8 O.Ord[T8], o9 O.Ord[T9]) O.Ord[Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
  return O.MakeOrd(func(l, r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    if c:= o4.Compare(l.F4, r.F4); c != 0 {return c}
    if c:= o5.Compare(l.F5, r.F5); c != 0 {return c}
    if c:= o6.Compare(l.F6, r.F6); c != 0 {return c}
    if c:= o7.Compare(l.F7, r.F7); c != 0 {return c}
    if c:= o8.Compare(l.F8, r.F8); c != 0 {return c}
    if c:= o9.Compare(l.F9, r.F9); c != 0 {return c}
    return 0
  }, func(l, r Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3) && o4.Equals(l.F4, r.F4) && o5.Equals(l.F5, r.F5) && o6.Equals(l.F6, r.F6) && o7.Equals(l.F7, r.F7) && o8.Equals(l.F8, r.F8) && o9.Equals(l.F9, r.F9)
  })
}

// Map9 maps each value of a [Tuple9] via a mapping function
func Map9[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, F4 ~func(T4) R4, F5 ~func(T5) R5, F6 ~func(T6) R6, F7 ~func(T7) R7, F8 ~func(T8) R8, F9 ~func(T9) R9, T1, R1, T2, R2, T3, R3, T4, R4, T5, R5, T6, R6, T7, R7, T8, R8, T9, R9 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9) func(Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Tuple9[R1, R2, R3, R4, R5, R6, R7, R8, R9] {
 return func(t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Tuple9[R1, R2, R3, R4, R5, R6, R7, R8, R9] {
    return MakeTuple9(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
      f4(t.F4),
      f5(t.F5),
      f6(t.F6),
      f7(t.F7),
      f8(t.F8),
      f9(t.F9),
    )
 }
}

// Replicate9 creates a [Tuple9] with all fields set to the input value `t`
func Replicate9[T any](t T) Tuple9[T, T, T, T, T, T, T, T, T] {
  return MakeTuple9(t, t, t, t, t, t, t, t, t)
}

// MakeTuple10 is a function that converts its 10 parameters into a [Tuple10]
func MakeTuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
  return Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10}
}

// Tupled10 converts a function with 10 parameters returning into a function taking a Tuple10
// The inverse function is [Untupled10]
func Tupled10[F ~func(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) R, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R any](f F) func(Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) R {
  return func(t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) R {
    return f(t.F1, t.F2, t.F3, t.F4, t.F5, t.F6, t.F7, t.F8, t.F9, t.F10)
  }
}

// Untupled10 converts a function with a [Tuple10] parameter into a function with 10 parameters
// The inverse function is [Tupled10]
func Untupled10[F ~func(Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) R, T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R any](f F) func(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) R {
  return func(t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, t9 T9, t10 T10) R {
    return f(MakeTuple10(t1, t2, t3, t4, t5, t6, t7, t8, t9, t10))
  }
}

// Monoid10 creates a [Monoid] for a [Tuple10] based on 10 monoids for the contained types
func Monoid10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](m1 M.Monoid[T1], m2 M.Monoid[T2], m3 M.Monoid[T3], m4 M.Monoid[T4], m5 M.Monoid[T5], m6 M.Monoid[T6], m7 M.Monoid[T7], m8 M.Monoid[T8], m9 M.Monoid[T9], m10 M.Monoid[T10]) M.Monoid[Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
  return M.MakeMonoid(func(l, r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
    return MakeTuple10(m1.Concat(l.F1, r.F1), m2.Concat(l.F2, r.F2), m3.Concat(l.F3, r.F3), m4.Concat(l.F4, r.F4), m5.Concat(l.F5, r.F5), m6.Concat(l.F6, r.F6), m7.Concat(l.F7, r.F7), m8.Concat(l.F8, r.F8), m9.Concat(l.F9, r.F9), m10.Concat(l.F10, r.F10))
  }, MakeTuple10(m1.Empty(), m2.Empty(), m3.Empty(), m4.Empty(), m5.Empty(), m6.Empty(), m7.Empty(), m8.Empty(), m9.Empty(), m10.Empty()))
}

// Ord10 creates n [Ord] for a [Tuple10] based on 10 [Ord]s for the contained types
func Ord10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](o1 O.Ord[T1], o2 O.Ord[T2], o3 O.Ord[T3], o4 O.Ord[T4], o5 O.Ord[T5], o6 O.Ord[T6], o7 O.Ord[T7], o8 O.Ord[T8], o9 O.Ord[T9], o10 O.Ord[T10]) O.Ord[Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]] {
  return O.MakeOrd(func(l, r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) int {
    if c:= o1.Compare(l.F1, r.F1); c != 0 {return c}
    if c:= o2.Compare(l.F2, r.F2); c != 0 {return c}
    if c:= o3.Compare(l.F3, r.F3); c != 0 {return c}
    if c:= o4.Compare(l.F4, r.F4); c != 0 {return c}
    if c:= o5.Compare(l.F5, r.F5); c != 0 {return c}
    if c:= o6.Compare(l.F6, r.F6); c != 0 {return c}
    if c:= o7.Compare(l.F7, r.F7); c != 0 {return c}
    if c:= o8.Compare(l.F8, r.F8); c != 0 {return c}
    if c:= o9.Compare(l.F9, r.F9); c != 0 {return c}
    if c:= o10.Compare(l.F10, r.F10); c != 0 {return c}
    return 0
  }, func(l, r Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) bool {
    return o1.Equals(l.F1, r.F1) && o2.Equals(l.F2, r.F2) && o3.Equals(l.F3, r.F3) && o4.Equals(l.F4, r.F4) && o5.Equals(l.F5, r.F5) && o6.Equals(l.F6, r.F6) && o7.Equals(l.F7, r.F7) && o8.Equals(l.F8, r.F8) && o9.Equals(l.F9, r.F9) && o10.Equals(l.F10, r.F10)
  })
}

// Map10 maps each value of a [Tuple10] via a mapping function
func Map10[F1 ~func(T1) R1, F2 ~func(T2) R2, F3 ~func(T3) R3, F4 ~func(T4) R4, F5 ~func(T5) R5, F6 ~func(T6) R6, F7 ~func(T7) R7, F8 ~func(T8) R8, F9 ~func(T9) R9, F10 ~func(T10) R10, T1, R1, T2, R2, T3, R3, T4, R4, T5, R5, T6, R6, T7, R7, T8, R8, T9, R9, T10, R10 any](f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6, f7 F7, f8 F8, f9 F9, f10 F10) func(Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Tuple10[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10] {
 return func(t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Tuple10[R1, R2, R3, R4, R5, R6, R7, R8, R9, R10] {
    return MakeTuple10(
      f1(t.F1),
      f2(t.F2),
      f3(t.F3),
      f4(t.F4),
      f5(t.F5),
      f6(t.F6),
      f7(t.F7),
      f8(t.F8),
      f9(t.F9),
      f10(t.F10),
    )
 }
}

// Replicate10 creates a [Tuple10] with all fields set to the input value `t`
func Replicate10[T any](t T) Tuple10[T, T, T, T, T, T, T, T, T, T] {
  return MakeTuple10(t, t, t, t, t, t, t, t, t, t)
}