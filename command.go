package main

//func (army *Army) Go(direction Direction) error {
//	for i := range *army {
//		switch (*army)[i].direction {
//		case North:
//			switch direction {
//			case Forward:
//				(*army)[i].x--
//			case Right:
//				(*army)[i].y++
//			case Backward:
//				(*army)[i].x++
//			case Left:
//				(*army)[i].y--
//			}
//		case East:
//			switch direction {
//			case Forward:
//				(*army)[i].y++
//			case Right:
//				(*army)[i].x++
//			case Backward:
//				(*army)[i].y--
//			case Left:
//				(*army)[i].x--
//			}
//		case South:
//			switch direction {
//			case Forward:
//				(*army)[i].x++
//			case Right:
//				(*army)[i].y--
//			case Backward:
//				(*army)[i].x--
//			case Left:
//				(*army)[i].y++
//			}
//		case West:
//			switch direction {
//			case Forward:
//				(*army)[i].y--
//			case Right:
//				(*army)[i].x--
//			case Backward:
//				(*army)[i].y++
//			case Left:
//				(*army)[i].x++
//			}
//		}
//	}
//	return nil
//}
