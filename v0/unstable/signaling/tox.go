package signaling

import (
	"fmt"
	"log"
	"strconv"

	toxcore "github.com/TokTok/go-toxcore-c"

	"github.com/punchguard/v0"
)

type Tox struct {
	punchguard.BaseSignaling
	tox *ToxIceTransport
}

func (t *Tox) Init(id string) error {
	// log.Printf("%T.Init", t)

	if err := t.InitBase(id); err != nil {
		return err
	}

	return nil
}

func toxInit(id string, t *Tox) *ToxIceTransport {
	var saveDataFilename, peerPublicKey string

	if id == "s1" {
		saveDataFilename = "./controlling.tox"
		peerPublicKey = "3F408CFCB0F4468ABD70EECC14AAA07A612FD382C6083A48D284CD944BE3677427A0BDD945BD"
	} else {
		saveDataFilename = "./controlled.tox"
		peerPublicKey = "D7AB298028495286F5A8CD233AD08CA51A95B24CFB780ECA412DA4B9E688861667B475FAC035"
	}

	return &ToxIceTransport{
		SaveDataFilename: saveDataFilename,
		PeerPublicKey:    peerPublicKey,
		t:                t,
	}
}

// curl -s https://nodes.tox.chat/json|jq ".nodes|.[]|[.ipv4,.port,.public_key]|tostring" -r|grep NONE -v|sed "s,],}\,,"|sed "s,\[,    {,"|pbcopy

func bootstrapNodes() [][3]string {
	return [][3]string{
		{"tox.initramfs.io", "33445", "3F0A45A268367C1BEA652F258C85F4A66DA76BCAA667A49E770BCC4917AB6A25"},
		{"144.217.167.73", "33445", "7E5668E0EE09E19F320AD47902419331FFEE147BB3606769CFBE921A2A2FD34C"},
		{"tox.abilinski.com", "33445", "10C00EB250C3233E343E2AEBA07115A5C28920E9C8D29492F6D00B29049EDC7E"},
		{"tox.novg.net", "33445", "D527E5847F8330D628DAB1814F0A422F6DC9D0A300E6C357634EE2DA88C35463"},
		{"198.199.98.108", "33445", "BEF0CFB37AF874BD17B9A8F9FE64C75521DB95A37D33C5BDB00E9CF58659C04F"},
		{"tox.kurnevsky.net", "33445", "82EF82BA33445A1F91A7DB27189ECFC0C013E06E3DA71F588ED692BED625EC23"},
		{"81.169.136.229", "33445", "E0DB78116AC6500398DDBA2AEEF3220BB116384CAB714C5D1FCD61EA2B69D75E"},
		{"tox2.abilinski.com", "33445", "7A6098B590BDC73F9723FC59F82B3F9085A64D1B213AAF8E610FD351930D052D"},
		{"46.101.197.175", "33445", "CD133B521159541FB1D326DE9850F5E56A6C724B5B8E5EB5CD8D950408E95707"},
		{"tox1.mf-net.eu", "33445", "B3E5FA80DC8EBD1149AD2AB35ED8B85BD546DEDE261CA593234C619249419506"},
		{"tox2.mf-net.eu", "33445", "70EA214FDE161E7432530605213F18F7427DC773E276B3E317A07531F548545F"},
		{"195.201.7.101", "33445", "B84E865125B4EC4C368CD047C72BCE447644A2DC31EF75BD2CDA345BFD310107"},
		{"tox4.plastiras.org", "33445", "836D1DA2BE12FE0E669334E437BE3FB02806F1528C2B2782113E0910C7711409"},
		{"gt.sot-te.ch", "33445", "F4F4856F1A311049E0262E9E0A160610284B434F46299988A9CB42BD3D494618"},
		{"188.225.9.167", "33445", "1911341A83E02503AB1FD6561BD64AF3A9D6C3F12B5FBB656976B2E678644A67"},
		{"122.116.39.151", "33445", "5716530A10D362867C8E87EE1CD5362A233BAFBBA4CF47FA73B7CAD368BD5E6E"},
		{"195.123.208.139", "33445", "534A589BA7427C631773D13083570F529238211893640C99D1507300F055FE73"},
		{"tox3.plastiras.org", "33445", "4B031C96673B6FF123269FF18F2847E1909A8A04642BBECD0189AC8AEEADAF64"},
		{"139.162.110.188", "33445", "F76A11284547163889DDC89A7738CF271797BF5E5E220643E97AD3C7E7903D55"},
		{"198.98.49.206", "33445", "28DB44A3CEEE69146469855DFFE5F54DA567F5D65E03EFB1D38BBAEFF2553255"},
		{"172.105.109.31", "33445", "D46E97CF995DC1820B92B7D899E152A217D36ABE22730FEA4B6BF1BFC06C617C"},
		{"91.146.66.26", "33445", "B5E7DAC610DBDE55F359C7F8690B294C8E4FCEC4385DE9525DBFA5523EAD9D53"},
		{"tox01.ky0uraku.xyz", "33445", "FD04EB03ABC5FC5266A93D37B4D6D6171C9931176DC68736629552D8EF0DE174"},
		{"tox02.ky0uraku.xyz", "33445", "D3D6D7C0C7009FC75406B0A49E475996C8C4F8BCE1E6FC5967DE427F8F600527"},
		{"tox.plastiras.org", "33445", "8E8B63299B3D520FB377FE5100E65E3322F7AE5B20A0ACED2981769FC5B43725"},
		{"kusoneko.moe", "33445", "BE7ED53CD924813507BA711FD40386062E6DC6F790EFA122C78F7CDEEE4B6D1B"},
		{"tox2.plastiras.org", "33445", "B6626D386BE7E3ACA107B46F48A5C4D522D29281750D44A0CBA6A2721E79C951"},
		{"172.104.215.182", "33445", "DA2BD927E01CD05EBCC2574EBE5BEBB10FF59AE0B2105A7D1E2B40E49BB20239"},
		{"91.219.59.156", "33445", "8E7D0B859922EF569298B4D261A8CCB5FEA14FB91ED412A7603A585A25698832"},
		{"85.143.221.42", "33445", "DA4E4ED4B697F2E9B000EEFE3A34B554ACD3F45F5C96EAEA2516DD7FF9AF7B43"},
		// {"tox.verdict.gg", "33445", "1C5293AEF2114717547B39DA8EA6F1E331E5E358B35F9B6B5F19317911C5F976"},
		{"78.46.73.141", "33445", "02807CF4F8BB8FB390CC3794BDF1E8449E9A8392C5D3F2200019DA9F1E812E46"},
		{"46.229.50.168", "33445", "813C8F4187833EF0655B10F7752141A352248462A567529A38B6BBF73E979307"},
		{"mk.tox.dcntrlzd.network", "33445", "5E815C25A4E58910A7350EC64ECB32BC9E1919F86844DC97125735C2C30FBE6E"},
		{"87.118.126.207", "33445", "0D303B1778CA102035DA01334E7B1855A45C3EFBC9A83B9D916FFDEBC6DD3B2E"},
		// {"floki.blog", "33445", "6C6AF2236F478F8305969CCFC7A7B67C6383558FF87716D38D55906E08E72667"},
		{"bg.tox.dcntrlzd.network", "33445", "20AD2A54D70E827302CDF5F11D7C43FA0EC987042C36628E64B2B721A1426E36"},
		{"46.146.229.184", "33445", "94750E94013586CCD989233A621747E2646F08F31102339452CADCF6DC2A760A"},
		{"209.59.144.175", "33445", "214B7FEA63227CAEC5BCBA87F7ABEEDB1A2FF6D18377DD86BF551B8E094D5F1E"},
		{"rs.tox.dcntrlzd.network", "33445", "FC4BADF62DCAF17168A4E3ACAD5D656CF424EDB5E0C0C2B9D77E509E74BD8F0D"},
		{"208.38.228.104", "33445", "3634666A51CA5BE1579C031BD31B20059280EB7C05406ED466BD9DFA53373271"},
		{"lunarfire.spdns.org", "33445", "E61F5963268A6306CCFE7AF98716345235763529957BD5F45889484654EE052B"},
		{"ru.tox.dcntrlzd.network", "33445", "DBB2E896990ECC383DA2E68A01CA148105E34F9B3B9356F2FE2B5096FDB62762"},
		{"43.231.185.239", "33445", "27D4029A96C9674C15B958011C62F63D4D35A23142EF2BA5CD9AF164162B3448"},
		{"141.95.108.234", "33445", "2DEF3156812324B1593A6442C937EAE0A8BD98DE529D2D4A7DD4BA6CB3ECF262"},
	}
}

const UINT32_MAX = 4294967295

type ToxIceTransport struct {
	punchguard.Async
	tox *toxcore.Tox

	peerId uint32

	SaveDataFilename string
	PeerPublicKey    string
	t                *Tox
}

func loadSaveData(saveDataFilename string) ([]byte, error) {
	if toxcore.FileExist(saveDataFilename) {
		data, err := toxcore.LoadSavedata(saveDataFilename)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, nil
}

func addFriend(tox *toxcore.Tox, peerPublicKey string) (uint32, error) {
	peer, err := tox.FriendByPublicKey(peerPublicKey)
	if err != nil {
		switch err.Error() {
		case fmt.Sprintf("toxcore error: %d", toxcore.ERR_FRIEND_BY_PUBLIC_KEY_NOT_FOUND):
			log.Printf("%T.addFriend: tox peer not found, adding %s", tox, peerPublicKey)
			peer, err = tox.FriendAddNorequest(peerPublicKey)
			if err != nil {
				return UINT32_MAX, err
			}
		default:
			return UINT32_MAX, err
		}
	}
	log.Printf("%T.addFriend: tox peer %s id %d", tox, peerPublicKey, peer)
	return peer, nil
}

func connectionAsText(connection int) string {
	switch connection {
	case toxcore.CONNECTION_NONE:
		return "Offline"
	case toxcore.CONNECTION_TCP:
		return "Online(TCP)"
	case toxcore.CONNECTION_UDP:
		return "Online(UDP)"
	default:
		return "UNKNOWN"
	}
}

func bootstrap(tox *toxcore.Tox) error {

	for _, node := range bootstrapNodes() {
		port, err := strconv.Atoi(node[1])
		if err != nil {
			return err
		}
		res, err := tox.Bootstrap(node[0], uint16(port), node[2])
		if err != nil {
			log.Printf("%T.bootstrap: %v %v %v", tox, res, err, node)
		}
	}
	log.Printf("%T.bootstrap: done", tox)
	return nil
}

func (t *ToxIceTransport) Open() error {

	// toxcore.SetDebug(true)
	// toxcore.SetLogLevel(toxcore.LOG_LEVEL_TRACE)
	// toxcore.SetLogLevel(toxcore.LOG_LEVEL_DEBUG)

	data, err := loadSaveData(t.SaveDataFilename)
	if err != nil {
		return err
	}

	options := toxcore.NewToxOptions()
	// options.Savedata_type = toxcore.SAVEDATA_TYPE_SECRET_KEY
	options.Savedata_type = toxcore.SAVEDATA_TYPE_TOX_SAVE
	options.Savedata_data = data
	options.Local_discovery_enabled = true
	options.LogCallback = func(tox *toxcore.Tox, level int, file string, line uint32, fname string, msg string) {
		log.Printf("%T.LogCallback %d %s %d %s %s", t, level, file, line, fname, msg)
	}

	tox := toxcore.NewTox(options)

	tox.CallbackSelfConnectionStatus(func(this *toxcore.Tox, status int, userData interface{}) {
		log.Printf("%T.CallbackSelfConnectionStatus: %s", t, connectionAsText(status))
	}, nil)

	tox.CallbackFriendConnectionStatus(func(tox *toxcore.Tox, friendNumber uint32, status int, userData interface{}) {
		log.Printf("%T.CallbackFriendConnectionStatus: %d %s", t, friendNumber, connectionAsText(status))

		if status == toxcore.CONNECTION_UDP {
			t.t.Ready()
		}
	}, nil)

	tox.CallbackFriendMessage(t.onCallbackFriendMessage, nil)

	t.peerId, err = addFriend(tox, t.PeerPublicKey)
	if err != nil {
		return err
	}

	err = bootstrap(tox)
	if err != nil {
		return err
	}

	tox.Bootstrap("tox.initramfs.io", 33445, "3F0A45A268367C1BEA652F258C85F4A66DA76BCAA667A49E770BCC4917AB6A25")

	log.Printf("%T.Open tox pub %s", t, tox.SelfGetPublicKey())
	log.Printf("%T.Open tox id  %s", t, tox.SelfGetAddress())
	log.Printf("%T.Open tox peer %s", t, t.PeerPublicKey)

	t.tox = tox

	return nil
}

func (t *ToxIceTransport) Close() error {
	err := t.tox.WriteSavedata(t.SaveDataFilename)
	t.tox.Kill()
	return err
}

func (t *ToxIceTransport) onCallbackFriendMessage(_ *toxcore.Tox, friendNumber uint32, message string, userData interface{}) {
	if t.peerId == friendNumber {
		// log.Printf("%T.onCallbackFriendMessage: %d %s", t, friendNumber, message)
		t.t.Receive(message)
	} else {
		log.Printf("%T.onCallbackFriendMessage: tox invalid friend: %d", t, friendNumber)
	}
}

func (m *ToxIceTransport) Send(msg string) {
	r, err := m.tox.FriendSendMessage(m.peerId, msg)
	// log.Printf("%T.FriendSendMessage: %v %v", m, msg, r)
	if err != nil {
		log.Printf("%T.Send: error %d %v", m, r, err)
	}
}

func (m *ToxIceTransport) Start() punchguard.StoppedEvent {
	var status int
	return m.MainLoop(func() {
		m.Open()
		status = m.tox.SelfGetConnectionStatus()
		log.Printf("%T.SelfGetConnectionStatus %s", m.tox, connectionAsText(status))
		log.Printf("%T.MainLoop: started", m)
	}, func() {
		if status != m.tox.SelfGetConnectionStatus() {
			status = m.tox.SelfGetConnectionStatus()
			log.Printf("%T.SelfGetConnectionStatus %s", m.tox, connectionAsText(status))
		}
		// pause := time.Duration(m.tox.IterationInterval())
		// time.Sleep(pause * time.Millisecond)
		m.tox.Iterate()
	}, func() {
		log.Printf("%T.MainLoop: stopped", m)
		m.Close()
	})
}

func (p *ToxIceTransport) Stop() {
	p.Async.StopAsync()
}

func (p *Tox) Stop() {
	if p.tox != nil {
		p.StopService(p.tox)
	}
	p.Async.StopAsync()
}

func (m *Tox) Start() punchguard.StoppedEvent {
	return m.MainLoop(func() {
		log.Printf("%T.MainLoop: started", m)
	}, func() {
		select {
		case <-m.OnConnectEvent():
			log.Printf("%T.OnConnectEvent", m)
			// TODO sync.RunOnce
			m.tox = toxInit(m.GetId(), m)
			m.tox.InitAsync()
			m.StartService(m.tox)
		case <-m.OnDisconnectEvent():
			log.Printf("%T.OnDisconnectEvent", m)
			if m.tox != nil {
				m.StopService(m.tox)
				m.tox = nil
			}
		case msg := <-m.OnSendEvent():
			m.WhenRunningAsync(func() {
				// log.Printf("%T.OnSendEvent '%v'", m, msg)
				if m.tox != nil {
					m.tox.Send(msg)
				}
			})
		default:
		}
	}, func() {
		log.Printf("%T.MainLoop: stopped", m)
		m.Close()
	})
}
