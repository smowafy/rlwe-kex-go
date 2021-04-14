package polynomial

import(
	"testing"
	"reflect"
)


var p1 []uint32 = []uint32 {
	324, 25, 306, 400, 92, 158, 152, 45, 50, 286, 105, 98, 64, 292, 151, 490, 149, 175, 229, 134, 17, 95,
	118, 124, 176, 303, 334, 26, 315, 315, 460, 478, 233, 68, 449, 282, 242, 77, 401, 399, 227, 444, 20, 27,
	116, 199, 20, 208, 191, 52, 319, 404, 400, 500, 206, 412, 335, 227, 300, 197, 448, 291, 226, 497, 422, 312,
	159, 23, 164, 472, 322, 340, 274, 356, 423, 222, 245, 157, 244, 342, 483, 231, 210, 300, 139, 435, 164, 171,
	204, 136, 190, 380, 475, 166, 4, 376, 417, 19, 245, 193, 317, 396, 55, 157, 438, 419, 389, 264, 156, 208,
	228, 443, 45, 4, 58, 12, 357, 332, 72, 89, 193, 187, 337, 430, 195, 70, 266, 216, 340, 490, 145, 335,
	226, 380, 23, 71, 0, 289, 448, 421, 77, 497, 221, 377, 140, 322, 20, 373, 254, 416, 89, 85, 30, 60,
	170, 170, 318, 274, 263, 291, 65, 15, 362, 160, 462, 239, 239, 346, 91, 492, 465, 369, 491, 440, 153, 382,
	225, 158, 486, 193, 463, 497, 310, 312, 368, 49, 236, 341, 487, 50, 209, 82, 370, 355, 151, 139, 103, 113,
	129, 448, 79, 192, 397, 398, 246, 356, 122, 327, 65, 37, 52, 477, 313, 186, 312, 111, 442, 267, 108, 93,
	280, 489, 375, 11, 349, 102, 486, 327, 8, 355, 121, 456, 365, 180, 361, 401, 472, 346, 474, 345, 315, 309,
	2, 388, 405, 20, 185, 24, 30, 253, 85, 133, 235, 36, 37, 151, 462, 265, 153, 317, 470, 100, 288, 73,
	19, 239, 250, 484, 234, 174, 201, 404, 369, 77, 149, 34, 32, 374, 455, 100, 9, 336, 93, 96, 414, 381,
	279, 11, 30, 256, 280, 468, 153, 206, 76, 200, 255, 106, 427, 139, 451, 491, 473, 1, 179, 436, 498, 42,
	4, 251, 265, 194, 387, 408, 245, 477, 469, 159, 426, 269, 248, 283, 231, 6, 143, 159, 323, 149, 146, 256,
	227, 455, 356, 432, 58, 127, 456, 393, 444, 323, 276, 53, 431, 9, 320, 341, 398, 334, 156, 23, 14, 109,
	117, 183, 148, 199, 328, 335, 385, 29, 250, 148, 185, 67, 222, 236, 172, 409, 444, 467, 214, 295, 438, 357,
	268, 500, 370, 155, 311, 90, 413, 148, 109, 362, 193, 141, 374, 75, 417, 233, 189, 376, 228, 199, 230, 82,
	71, 115, 468, 412, 105, 436, 114, 324, 47, 21, 77, 396, 269, 452, 191, 191, 104, 120, 498, 290, 190, 53,
	9, 267, 348, 168, 297, 129, 308, 500, 242, 488, 474, 195, 32, 199, 427, 351, 404, 418, 132, 344, 467, 422,
	263, 300, 34, 388, 20, 232, 437, 132, 475, 55, 162, 215, 380, 446, 347, 493, 459, 45, 363, 292, 316, 285,
	486, 385, 337, 402, 153, 342, 302, 428, 57, 372, 22, 416, 424, 219, 210, 117, 387, 270, 310, 347, 230, 407,
	71, 176, 121, 208, 306, 464, 125, 109, 72, 450, 494, 360, 60, 257, 182, 198, 377, 181, 336, 345, 32, 482,
	350, 306, 227, 409, 85, 0,
}

var p2 []uint32 = []uint32 {
	330, 8, 230, 84, 315, 383, 88, 47, 42, 242, 253, 31, 492, 348, 440, 217, 340, 139, 139, 54, 392, 304,
	481, 455, 449, 261, 143, 25, 396, 75, 89, 292, 433, 351, 39, 117, 262, 426, 107, 491, 167, 435, 492, 256,
	274, 352, 138, 484, 404, 119, 489, 282, 60, 112, 80, 362, 182, 91, 198, 142, 24, 105, 255, 423, 202, 296,
	405, 87, 297, 396, 127, 154, 456, 191, 232, 31, 476, 55, 124, 65, 78, 417, 111, 86, 338, 403, 75, 248,
	478, 362, 365, 247, 443, 269, 256, 267, 324, 105, 484, 48, 282, 394, 446, 293, 27, 223, 204, 243, 406, 190,
	30, 411, 231, 313, 146, 3, 467, 500, 45, 169, 218, 162, 23, 93, 411, 377, 184, 381, 368, 327, 304, 162,
	296, 126, 108, 367, 148, 84, 20, 330, 218, 306, 32, 116, 461, 135, 294, 239, 310, 340, 45, 22, 26, 133,
	307, 499, 244, 374, 486, 259, 91, 2, 457, 124, 467, 449, 197, 477, 450, 3, 170, 186, 153, 296, 358, 209,
	170, 61, 149, 207, 137, 438, 462, 498, 314, 98, 385, 172, 446, 457, 397, 274, 40, 121, 33, 235, 489, 239,
	39, 7, 256, 48, 233, 8, 325, 13, 200, 61, 220, 420, 396, 269, 263, 445, 421, 278, 102, 229, 134, 351,
	300, 333, 129, 474, 212, 405, 297, 141, 312, 224, 106, 466, 439, 352, 6, 298, 6, 36, 447, 107, 32, 111,
	492, 14, 188, 196, 220, 118, 381, 252, 214, 152, 340, 322, 205, 230, 68, 412, 425, 309, 209, 40, 90, 119,
	329, 386, 269, 354, 353, 467, 419, 3, 468, 222, 252, 391, 383, 208, 247, 347, 145, 471, 263, 293, 254, 438,
	198, 214, 103, 356, 363, 354, 402, 334, 201, 223, 38, 177, 235, 7, 20, 303, 43, 170, 171, 250, 143, 389,
	495, 262, 337, 360, 145, 238, 300, 186, 153, 448, 102, 453, 175, 97, 385, 11, 69, 191, 489, 12, 128, 239,
	444, 240, 9, 269, 166, 46, 340, 333, 475, 226, 167, 327, 421, 299, 446, 56, 39, 440, 346, 448, 150, 486,
	59, 339, 300, 288, 500, 438, 156, 9, 193, 218, 251, 172, 174, 395, 99, 336, 151, 121, 334, 103, 183, 407,
	204, 353, 102, 177, 104, 466, 57, 224, 449, 108, 106, 412, 58, 450, 117, 452, 148, 66, 146, 79, 279, 357,
	107, 336, 440, 232, 220, 203, 86, 196, 406, 299, 443, 475, 198, 378, 359, 296, 291, 465, 347, 74, 313, 477,
	491, 438, 188, 199, 360, 336, 13, 360, 7, 53, 374, 72, 80, 308, 452, 344, 437, 414, 453, 365, 100, 69,
	481, 59, 354, 306, 151, 302, 21, 328, 395, 349, 86, 349, 247, 399, 236, 467, 292, 94, 481, 403, 24, 116,
	92, 101, 266, 364, 70, 92, 130, 431, 213, 317, 449, 102, 26, 68, 149, 492, 442, 125, 457, 314, 240, 74,
	297, 29, 42, 292, 300, 285, 87, 331, 134, 374, 183, 374, 138, 188, 449, 25, 436, 280, 96, 390, 175, 156,
	416, 337, 352, 466, 283, 0,
}

var p1p2 []uint32 = []uint32 {
	4262829072, 4263129822, 4263381319, 4263204700, 4263314628, 4264005662, 4263478320, 4263272877, 4263542132, 4263784086, 4262984694, 4263753584, 4263317310, 4263985592, 4263878710, 4263841443, 4263515704, 4263908647, 4264460349, 4265137409, 4264847058, 4264731926,
	4265306232, 4265088832, 4265010172, 4265483824, 4265806836, 4265602156, 4265929672, 4265138373, 4266066316, 4265760570, 4265296516, 4266651222, 4266156563, 4266955118, 4265908811, 4267309697, 4266558184, 4266541646, 4267947321, 4267456957, 4266727489, 4268064779,
	4267363416, 4267324549, 4268146423, 4268200268, 4267967098, 4268766631, 4269253458, 4268040387, 4268776165, 4268935369, 4269912703, 4269638207, 4269767331, 4268744713, 4269411163, 4269572521, 4270306748, 4269807895, 4270393209, 4270126011, 4270172453, 4270842827,
	4270184011, 4270538291, 4270937647, 4269903227, 4269768854, 4271756012, 4271752080, 4272024262, 4271617690, 4271358381, 4272167430, 4271838703, 4272018943, 4272138916, 4272171439, 4273022960, 4271995347, 4273071264, 4272812826, 4272883182, 4273019114, 4273276785,
	4272927956, 4273623137, 4274238136, 4273369511, 4274720625, 4274259642, 4274554061, 4275076996, 4275069848, 4274797700, 4274185142, 4273724415, 4276457925, 4275555163, 4275785255, 4276101523, 4276016848, 4276324446, 4276083394, 4276599776, 4275395145, 4275114444,
	4275828808, 4276434952, 4275960461, 4277314421, 4276445261, 4276832933, 4277331088, 4276805131, 4277386708, 4278495323, 4278250953, 4277743889, 4277328133, 4277418835, 4277702731, 4277659992, 4278117422, 4277747682, 4279107684, 4278353137, 4277999195, 4278718955,
	4279202517, 4278682002, 4278566133, 4277547583, 4278356982, 4278725867, 4278708758, 4278928932, 4278761305, 4279591626, 4280563552, 4280774361, 4279244805, 4279531198, 4280340301, 4280922601, 4280888255, 4281205172, 4280471998, 4281012221, 4280468514, 4281311995,
	4280851287, 4281204262, 4281369659, 4280908707, 4281884390, 4281634422, 4281589074, 4281972024, 4282001604, 4281352668, 4282227282, 4282140753, 4282427642, 4282175489, 4282504114, 4283195009, 4284085646, 4282913877, 4282993195, 4283510536, 4283835159, 4283598648,
	4283670072, 4284209701, 4283587968, 4283766592, 4285956588, 4284602596, 4284768520, 4285338611, 4285473789, 4286169744, 4286477561, 4286324751, 4286717464, 4286935221, 4286083382, 4286143349, 4287042699, 4287213908, 4287150380, 4286770967, 4287632719, 4286631001,
	4287466737, 4286739935, 4286509221, 4286410928, 4286815738, 4287037172, 4288047982, 4287404983, 4286955045, 4287446000, 4286770026, 4288019834, 4287433851, 4288511020, 4288912308, 4288547023, 4288731117, 4288871904, 4288785786, 4289563126, 4289397195, 4289496000,
	4289718168, 4289425464, 4289839994, 4289933756, 4289496249, 4289838833, 4290335903, 4290752196, 4290649751, 4290604699, 4290708406, 4291007782, 4290159710, 4290592128, 4291654252, 4291451724, 4292418094, 4290901257, 4291815061, 4292325362, 4292970943, 4292304404,
	4291603010, 4292619440, 4293438847, 4293343770, 4292686937, 4293931017, 4293626314, 4293951971, 4293824441, 4292562116, 4293198063, 4292994575, 4294369870, 4293723439, 4293816614, 4294048480, 4293609571, 4294198561, 4294527268, 4293982763, 4294351193, 4294234487,
	4294515396, 4294839707, 4294554679, 4294874810, 530659, 10303, 4294905220, 1493694, 907136, 1519962, 811434, 783058, 1368073, 849448, 786114, 1732841, 1709331, 1721948, 2323342, 2814710, 2625652, 2149020,
	2731343, 2952958, 3145821, 3148547, 2705498, 3011857, 2647090, 3561353, 3889709, 3509394, 3095243, 2962513, 3328582, 3019686, 4255833, 3664055, 3868674, 4413068, 4196607, 5725859, 5494329, 5003654,
	3931656, 5244733, 5414422, 5567960, 5579146, 5784970, 5626336, 6006718, 6259302, 6666863, 6641307, 6687482, 7203156, 6541616, 7965621, 6965757, 7461308, 6833115, 6741622, 7790281, 8729919, 8145068,
	7614849, 7789824, 7942059, 8137402, 7814086, 7465030, 8640289, 9701641, 8932202, 9166386, 9157650, 9416966, 9859806, 10564392, 8824029, 10278562, 10649130, 10746721, 10434191, 10883125, 9724809, 10869493,
	11442084, 10882548, 11353090, 11035046, 10813807, 11484850, 11288395, 11370245, 12509662, 12026039, 11853132, 10947944, 11730179, 11846898, 11484164, 11847437, 12559540, 12699023, 13945675, 12376898, 12500172, 13264851,
	12625487, 12693528, 13378008, 13703318, 13663999, 14335078, 13738483, 14914996, 14209370, 14884773, 13985014, 14842204, 14827992, 15163120, 15814649, 15529304, 15862245, 15237571, 15339125, 15445299, 15960630, 15717966,
	16298018, 16286377, 16415804, 15707513, 16488529, 17160664, 17364313, 17011787, 16429729, 16396948, 16938185, 18243974, 17777602, 17586074, 17150532, 18384788, 18737622, 19222019, 19114644, 18333667, 18433305, 19380201,
	18861743, 18618129, 19147877, 19345388, 19683219, 19865985, 19378977, 20043856, 20441445, 19872046, 19527246, 19604429, 21151534, 20148549, 20639509, 20699391, 21772346, 21250374, 21790417, 21686647, 22113443, 22554880,
	22968317, 21823793, 22552207, 22246197, 21979325, 24204189, 23004403, 23840194, 24719387, 24227533, 23598890, 23379121, 23916012, 24575293, 23771282, 24419792, 25275997, 25485417, 25832461, 25222380, 25890814, 25128940,
	25541181, 25546074, 26784318, 25612572, 25862173, 27177498, 26443405, 27066809, 27055460, 27162783, 26864702, 27922299, 27926939, 27215710, 27038913, 27584014, 27678573, 28363503, 28990656, 28816468, 28053914, 29087523,
	28602899, 28878961, 29668858, 29281213, 29494559, 28969144, 29195347, 29226374, 29203873, 30138711, 29873494, 30616262, 30454278, 30246032, 30103272, 30017295, 30580020, 30871692, 30908036, 31507574, 31360596, 31171003,
	31796665, 30463067, 31414638, 32347108, 32953007, 32220821,
}

func TestNussbaumerPolynomial(t *testing.T) {
	acoefs := []uint32 { 1, 3, 2, 5, }

	expected := [][]uint32 { []uint32{1, 2}, []uint32{3, 5} }

	res := NussbaumerPolynomial(acoefs)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("nussbaumer polynomial failed, expected %v, got %v\n", expected, res)
	}
}

func TestNaiveMultiply(t *testing.T) {
	res := NaiveMultiply(p1, p2)

	if !reflect.DeepEqual(res, p1p2) {
		t.Errorf("naive multiply failed, expected %v, got %v\n", p1p2, res)
	}
}

func TestNussbaumerIterativeMultiply(t *testing.T) {

	res := NussbaumerIterativeMultiply(p1, p2)

	if !reflect.DeepEqual(res, p1p2) {
		t.Errorf("nussbaumer multiply failed, expected:\n%v\n\ngot %v\n", p1p2, res)
	}
}

func BenchmarkNaiveMultiply(b *testing.B) {
	a0 := make([]uint32, 512)
	a1 := make([]uint32, 512)

	for i := 0; i < len(a0); i++ {
		a0[i] = EnsureMod(RandomUInt32(NewRandomGenerator()))
		a1[i] = EnsureMod(RandomUInt32(NewRandomGenerator()))
	}

	b.ResetTimer()

	NaiveMultiply(a0, a1)
}

func BenchmarkNussbaumerIterativeMultiply(b *testing.B) {
	a0 := make([]uint32, 512)
	a1 := make([]uint32, 512)

	for i := 0; i < len(a0); i++ {
		a0[i] = EnsureMod(RandomUInt32(NewRandomGenerator()))
		a1[i] = EnsureMod(RandomUInt32(NewRandomGenerator()))
	}

	b.ResetTimer()

	NussbaumerIterativeMultiply(a0, a1)
}

func BenchmarkNussbaumerRecursiveMultiply(b *testing.B) {
	a0 := make([]uint32, 512)
	a1 := make([]uint32, 512)

	for i := 0; i < len(a0); i++ {
		a0[i] = EnsureMod(RandomUInt32(NewRandomGenerator()))
		a1[i] = EnsureMod(RandomUInt32(NewRandomGenerator()))
	}

	b.ResetTimer()

	NussbaumerRecursiveMultiply(a0, a1)
}
