// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chacha20poly1305

var chacha20Poly1305IETFTests = []testVector{
	{
		"808182838485868788898a8b8c8d8e8f909192939495969798999a9b9c9d9e9f",

		"070000004041424344454647",

		"50515253c0c1c2c3c4c5c6c7",

		"4c616469657320616e642047656e746c656d656e206f662074686520636c617373206f66202739393a204966204920636f756c64206f6666657220796f75206f6e6c79206f6e652074697020666f7220746865206675747572652c2073756e73637265656e20776f756c642062652069742e",

		"d31a8d34648e60db7b86afbc53ef7ec2a4aded51296e08fea9e2b5a736ee62d63dbea45e8ca9671282fafb69da92728b1a71de0a9e060b2905d6a5b67ecd3b3692ddbd7f2d778b8c9803aee328091b58fab324e4fad675945585808b4831d7bc3ff4def08e4b7a9de576d26586cec64b61161ae10b594f09e26a7e902ecbd0600691",
	},
	{
		"a5117e70953568bf750862df9e6f92af81677c3a188e847917a4a915bda7792e",

		"129039b5572e8a7a8131f76a",

		"00000000000000001603030010",

		"1400000cebccee3bf561b292340fec60",

		"2b487a2941bc07f3cc76d1a531662588ee7c2598e59778c24d5b27559a80d163",
	},
	{
		"a5117e70953568bf750862df9e6f92af81677c3a188e847917a4a915bda7792e",

		"129039b5572e8a7a8131f76a",

		"00000000000000000000000000",

		"0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",

		"3f487a25aa70e9c8391763370569c9e83b7650dd1921c8b78869f241f25d2096c910b180930c5b8747fd90959fe8ca2dcadb4fa50fa1439f916b2301e1cc0810d6725775d3ab86721700f96e22709b0a7a8bef32627dd929b2dd3ba15772b669062bb558bc92e6c241a1d60d9f0035e80c335f854815fe1138ab8af653eab3e122135feeec7dfaba1cc24af82a2b7acccdd824899a7e03cc29c25be8a4f56a66673845b93bae1556f09dafc89a0d22af207718e2a6bb022e9d917597295992ea3b750cc0e7a7c3d33b23c5a8aeab45f5bb542f6c9e6c1747ae5a344aff483ba38577ad534b33b3abc7d284776ea33ed488c2a2475648a4fcda561745ea7787ed60f2368deb27c75adce6ff9b6cc6de1f5e72a741e2d59f64751b3ae482d714e0c90e83c671ff98ed611823afb39e6e5019a6ba548a2a72e829c7b7b4a101ac9deb90a25d3e0c50d22e1fc26c7c02296fa13c6d9c14767f68aaf46450a8d0fd5feb60d9d73c6e68623425b4984a79d619dd6bf896459aa77a681ec9c1a97f645e121f47779b051f8948a817f84d1f55da170d5bbbaf2f64e18b97ed3fd822db2819f523314f1e5ac72e8f69bbe6c87c22daddb0e1ac6790f8534071de2f258064b99789bfb165b065b8fe96f9127cd7dca9f7cb0368420f1e802faa3ca23792f2a5b93773dd405e71c320b211b54f7a26626b03c060e1ab87f32ac588abfa056ce090bd7c69913a700c80f325bfe824fa",
	},
	{
		"f2db28620582e05f00f31c808475ca3df1c20e340bf14828352499466d79295f",

		"4349e2131d44dc711148dfe3",

		"e4f0a3a4f90a8250f8806aa319053e8d73c62f150e2f239563037e9cc92823ad18c65111d0d462c954cc6c6ed2aafb45702a5a7e597d13bd8091594ab97cf7d1",

		"0967de57eefe1aaa999b9b746d88a1a248000d8734e0e938c6aa87",

		"bd06cc144fdc0d8b735fa4452eabbf78fd4ad2966ea41a84f68da40ca2da439777bc2ba6c4ec2de0d003eb",
	},
	{
		"05ff881d1e151bab4ca3db7d44880222733fe62686f71ce1e4610f2ea19599a7",

		"b34710f65aed442e4a40866b",

		"85deea3dc4",

		"c4c920fb52a56fe66eaa8aa3fa187c543e3db8e5c8094c4313dc4ed35dfc5821c5791d171e8cfe8d37883031a0ad",

		"b154452fb7e85d175dd0b0db08591565c5587a725cf22386922f5d27a01015aba778975510b38754b2182e24352f019b7ad493e1ed255906715644aec6e0",
	},
	{
		"dd0f2d4bb1c9e5ca5aa5f38d69bc8402f7dbb7229857b4a41b3044d481b7655e",

		"2bbca0910cc47ca0b8517391",

		"72611bef65eb664f24ea94f4d5d3d88c9c9c6da29c9a1991c02833c4c9f6993b57b5",

		"c4b337df5e83823900c6c202e93541cf5bc8c677a9aad8b8d87a4d7221e294e595cbc4f34e462d4e0def50f62491c57f598cf60236cfba0f4908816aea154f80e013732e59a07c668fcc5cb35d2232b7ae29b9e4f874f3417c74ab6689fae6690d5a9766fa13cd8adf293d3d4b70f4f999adde9121d1d29d467d04cf77ea398444d0ea3fe4b7c9c3e106002c76f4260fa204a0c3d5",

		"83aa28d6d98901e2981d21d3758ae4db8cce07fe08d82ca6f036a68daa88a7dda56eeb38040c942bdda0fd2d369eec44bd070e2c9314992f68dc16989a6ac0c3912c378cf3254f4bae74a66b075e828df6f855c0d8a827ffed3c03582c12a9112eeb7be43dfe8bd78beb2d1e56678b99a0372531727cb7f2b98d2f917ec10de93fe86267100c20356e80528c5066688c8b7acba76e591449952343f663993d5b642e59eb0f",
	},
	{
		"9cf77bd06a4ed8fb59349791b98ba40b6019611942f5768e8be2ee88477149e3",

		"b928935c4c966c60fd6583c0",

		"74ba3372d308910b5c9c3885f41252d57556",

		"a9775b8e42b63335439cf1c79fe8a3560b3baebfdfc9ef239d70da02cea0947817f00659a63a8ee9d67fb1756854cc738f7a326e432191e1916be35f0b78d72268de7c0e180af7ee8aa864f2fc30658baa97f9edb88ace49f5b2a8002a8023925e9fa076a997643340c8253cf88ac8a221c190d94c5e224110cb423a4b65cca9046c1fad0483e1444c0680449148e7b20a778c56d5ae97e679d920c43eed6d42598cf05d10d1a15cd722a0686a871b74fea7cad45562bacf3bda937ac701bc218dac7e9d7d20f955429abdac21d821207febf4d54daea4898837035038bf71c66cef63e90f5d3e51f7fcfe18d41f38540a2c2958dacde16304e4b33da324030f1366f923c337",

		"ec7fd64fd75b254961a2b7fc942470d8620f439258b871d0d00f58028b5e0bee5e139e8108ac439391465d6658f559b1df57aa21cf826ede1a28bc11af885e13eebfc009870928fae8abfdd943a60c54fca93f0502dc23d29c2fd5340f9bc0e6ef2a18b66ef627af95f796d5bbca50de22c8ec802da9397089b25c6ba5262468e3977b45dc112e51896c70731b0a52d7efec7c93b41995823436bf4b0c477ae79684407c9831b487928b2b8303caca752b3edf1f0598e15831155462706f94ef3fa3a9e5f937f37085afa9b4bbf939d275796a61b78f70597acfd25cd87f967021cd99328fc371b5eb5739869520657b30e4a5b0db7c8715cbe275dee78e719b357d3a9731f9eaba95986479bb2004a77822fc115a3d",
	},
	{
		"7ee32dd501dce849cd492f6e23324c1a4567bfceff9f11d1352bcb8615f1b093",

		"8998e043d2961afa51ea262a",

		"7e8da4f3018f673f8e43bd7a1dee05f8031ec49129c361abbc2a434e9eaf791c3c1d0f3dad767d3bba3ab6d728bbcf2bd994bd03571eae1348f161e6a1da03ddf7121ba4",

		"b3d3128bce6bbf66fd78f1a18352bae56bfcdae18b65c379ee0aeb37ee54fba1270d2df578ec5b75654d16e89fd1cd0acda7ec580dafd2fbbabd32a8112d49383a762db2638928c8d63eb0750f7e7fdd256b35321b072dd5c45f7dd58cc60dc63d3b79a0c4a1689adf180fef968eccbcfa01ee15091ceacd7b67a3082db0ce6aeb470aafe87249c88b58b721e783dde184ccf68de8e05b6347fe6b74ae3adf9a81e9496a5c9332e7ebe908d26ce6b3f0b2a97e9a89d9fdd0d7694585a3241f240d698e69fcc050e7a959ba153f6d06f117848ba05d887134f1b6b994dad9b9e74247513e08a125b1fadfc7394dcd2a6451b504ae3e75e22f2b9bc405747dedb6c43ef4ccdf1a7edaf9451346123eaa63f3af113124f361508e255503a242b96680ae3360c8b13ac1f64d08088bb26b7f617cb0866f11d6fd362b00d86eba3fee68724e302388f119d6f92161ac8ce00d08919377a26974d99575b1032ff0f1976240c785c8b89e9eb2bf005e4be06b5371ffca14683fedfdb49e00e38ff27af1324177faf91599abd5990920797574eb743effdc7decda318ada1419cc8e0bfecf82f9c99792746c2b",

		"ba85e72af18cb5ba85a4a0d6c28b4ac1e5509a3a2fdb0e3255cbc559df5e6a661fc560c756a0264dd99b72c61c51a4b7ad56ca4c8ccb7e8edfc48ff3cceac5d1e8ac5fc87096adc4d0e9a27492857b17604c3a694cfe0e70b22df106c8f3c61f840bcd634964cdb571840e125e381e7dd3a0d97972e965f16f775fa4ce555124318290bf508beb7bd77e633042deb0e863631478fc3dc9122862b3c31264471bcce54e0b74040c8bafd481cf798f332e8940f1134d3027d6f28e771d15e154fc89c6c25fe18a5d312807cc2e623bb1bbb4f0b6ec71d009407eb54bb0759f03682f65d0da8812f84d8e97483f6a8d76a8417efcd9526444abba24288647609791578887ef49780b0b89f51b072cae81c5b5014463da3633dda105b82add0f9c2f065dca46eedd2928be2570493c79a996fa78ea6aec0996497fe2dc444432ade4eaa662ee2255f0f4b92d593288a8e3ffe7a15a10e9d33b0203af23f4c9fd2cfcb6160db63b52810869ff1e65423dbe2c4415884b9f8dec3c968e14cd74f323c89053a96111bc9ce59ec483832c49c53a648e5f0f797f53642ac60170c94b473f1f2e7d8a38e46460b81219b52081263027f74cbf63a75af3a7",
	},
	{
		"7d5c4314a542aff57a454b274a7999dfdc5f878a159c29be27dabdfcf7c06975",

		"aeb6159fa88bb1ffd51d036d",

		"63b90dd89066ad7b61cc39497899a8f14399eace1810f5fe3b76d2501f5d8f83169c5ba602082164d45aad4df3553e36ef29050739fa067470d8c58f3554124bf06df1f27612564a6c04976059d69648ff9b50389556ad052e729563c6a7",

		"68d5ba501e87994ef6bc8042d7c5a99693a835a4796ad044f0e536a0790a7ee1e03832fec0cb4cb688cdf85f92a1f526492acac2949a0684803c24f947a3da27db0c259bd87251603f49bfd1eab4f733dec2f5725cfcf6dc381ad57fbdb0a699bccc34943e86f47dcfb34eba6746ed4508e3b764dfad4117c8169785c63d1e8309531747d90cc4a8bf13622759506c613324c512d10629991dc01fe3fe3d6607907e4f698a1312492674707fc4dde0f701a609d2ac336cc9f38badf1c813f9599148c21b5bd4658249d5010db2e205b3880e863441f2fe357dab2645be1f9e5067616bc335d0457ea6468c5828910cb09f92e5e184e316018e3c464c5ce59cc34608867bd8cbfa7e1286d73a17e3ebb675d097f9b3adfa41ea408d46252a096b3290e70a5be1896d6760a87e439334b863ccb11679ab5763ebe4a9110eb37c4043634b9e44d40cab34b42977475e2faa2ae0c0a38b170776fbb0870a63044aa6679545ac6951579d0581144cdf43f60923b6acaecdb325c864acd2c7b01d6e18b2b3c41c041bb9099cce557b114b84350131e3cee4089648b5691065867e7d38314154355d0e3ef9dc9375eddef922df2a06ad0f0e4357c3ac672932e5a66b16e8bf4b45cd893ea91cb397faadb9d9d7bf86e6ceca3e9176a5baa98b6114a149d3ed8ea176cc4a9380e18d2d9b67045aedeb28b729ba2ece74d759d5ebfb1ebee8ac5f5e79aaf1f98b7f2626e62a81d315a98b3e",

		"7597f7f44191e815a409754db7fea688e0105c987fa065e621823ea6dea617aed613092ad566c487cfa1a93f556615d2a575fb30ac34b11e19cd908d74545906f929dc9e59f6f1e1e6eaaabe182748ef87057ef7820ffcf254c40237d3ea9ff004472db783ed54b5a294a46cf90519bf89367b04fc01ce544c5bcdd3197eb1237923ce2c0c99921ca959c53b54176d292e97f6d9696ded6054711721aebda543e3e077c90e6f216cdc275b86d45603521c5aab24f08fd06833b0743c388382f941e19e0283ac7c4ef22383e1b9b08572882769c1382bab9ad127e7f3e09b5330b82d3e0c7d6f0df46edc93265999eef8e7afa0cb1db77df7accf5bff8631a320d146a5c751a637a80f627b0c9a41b44f09212f38c154226de02f4906ef34139bbeacc3f06739c8540e37334392d38ba1cbf4bc7debe77c09b35d2200216db15ed4389f43bfd8ae9bf76fd8243c3d869546e16b8e44a6cd1edbd2c58ef890b5a84cda889131e5cd9402ca4d8271052c6b4fe3f2dff54fb77bcb575c315b9109f90b14bc8e109919808a581c1809e2a188d29fd34ce639088a6683f641925f5b4b3529baa34e080bb47fb7ad9b43d0d67c9e6ae7cacb50527fa74e56d0c8b20149f5d332d686d48ebbe634c2b5d35fc84c69a5bcc93b93dedcf9fdf19a1fb9b75f6df9692d16f6c3490377a06294499e4b8ebeaa0cfd840bfa05fde21c0b5e94d13063b3f5da7b537caefe89069cfa9de9eb8f06e4d30125de64716f821bcc8279c0c7ea2e",
	},
	{
		"a5117e70953568bf750862df9e6f92af81677c3a188e847917a4a915bda7792e",

		"129039b5572e8a7a8131f76a",

		"73afd2ab0e0e8537cae42dc6530dc4afb6934ca6",

		"135a28170fe89066da7bcff3a9ccc1b27dfe942a6f47b23835ef746aaea63dc10066d90f4e697528e5451b8e11dd408fdbd4b94a1c6c82515bf7bc099df9cb9d5fa4acad0d22d5f267f18078cec107a995c1f3b12d7603886dbf910ab85ca7180053c50e759b00dc8c81555a425c03d71df6894a6c8cd2d94b64e303c08a1bc1dee1cf537ccf300850856292e1656aff5bf349c87f1ca1ca8085cd400fe901edcad04146a0714ef0f6b083d715edd670e020385f3cda29bc5ff6fc6edffe5ca9ce9def6e0e3d5f04ede2db02cfb2",

		"2c125232a59879aee36cacc4aca5085a4688c4f776667a8fbd86862b5cfb1d57c976688fdd652eafa2b88b1b8e358aa2110ff6ef13cdc1ceca9c9f087c35c38d89d6fbd8de89538070f17916ecb19ca3ef4a1c834f0bdaa1df62aaabef2e117106787056c909e61ecd208357dd5c363f11c5d6cf24992cc873cf69f59360a820fcf290bd90b2cab24c47286acb4e1033962b6d41e562a206a94796a8ab1c6b8bade804ff9bdf5ba6062d2c1f8fe0f4dfc05720bd9a612b92c26789f9f6a7ce43f5e8e3aee99a9cd7d6c11eaa611983c36935b0dda57d898a60a0ab7c4b54",
	},
}