package core

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrAccountNotFound     = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient account balance")
)

type Account struct {
	Address string
	Balance uint64
}

func (a *Account) String() string {
	return fmt.Sprintf("%d", a.Balance)
}



type AccountState struct {
	mu       sync.RWMutex
	accounts map[string]*Account
}

func NewAccountState() *AccountState {
	return &AccountState{
		accounts: map[string]*Account{
			"1ba2d5da9c9f12281cfc0aad9c9782d825f8154d5a86eff1bc37b49f56d042d157c6770e8cb62d0de2b1d9932b1dda55bd66b651cf14bfd1d609bb804f458cb95bd62792eb0d219fa35bcc9ba10dc15fec8768c8894e100e7fabbf608e91649ee4107403d12f3e34d0bb6dfbda0d70a459d24c1b1c762228836382d703997fcabdac9fc7021d548ff01c8ed49def248034f29ef66f3a067ee8c485ade8b1369ac215f893d14bbea6729a821ebaea8467dcd10d752c0a91aadedd844baa4337da977d13ce5d52a5c98ac223a46117dc08c2518a39d07c6d5608cbbe3412f7abc142b164f7de702d8a24e11b4ecc4147d2de4b3f09d46893c3adca7aaf230f9649536cad7547592bf126e1b74f455fc1f5ac7b45da864f9b15bb98170d276fe8e293f085cac8ce4378e30ed2ee06e1acb4b6b9295073239ff7a9cd6df354fdf9c81e6343f0adc5b7dda6c62659bdca04948eda992cb62d4373ae3ae81f7d6bf687caafdf85b284398c3f615bd4967070a6ab7ad08a80ec3428d62471b13cfcac88e7afb4166261fc5fd488dcc6ede645ca4c627e898568ad4c471146ea8681540e0b0cfde95029e57ba0acbc3e05951130733d1d1d19a69b6357c44e50e2f0c9dbb841afc6b29a5f0051cbc9af0f312c6405fa726290f658a8be4b4aa8f3f4e23cdb91926b14068ed544231f798f38b1b6f1a94d2d10785cc179bb858a0ab802aad0f4d15ce47c79d7faf3c5841f3126b3fd0f42f1ba5f386eaefa1d8fe637937cbeba9012856428e721aaca67203b44b51f5b9f695fdd6cf98f2b49b26e194776347866460da601eb8ae738afbcc1268042386df5384aa62b8c64a46a27fc58e4c6e03adeef56f5629f2e9df42b20cd0f3dae2a9eb07c38ede7fac9c5573deae9594b1e03760994c6cba502af2664f2d85b1ee24a9ee952c8eacb99d7d7110ce8c06704af474d11c5c0adf95191ba44888b46afa7b8927f4b21f57796de709d48f184694b6e3ce43c6525cf3f37d72b2db1b0910e9c8feaaa553f49b844a2af255018f0a8b877ac9c03f40206d9a3d2f8506bc2666c05246fdc10b65fde588ca06d86d96cc5e8c567ae0cd34a159506c35fe0839acb856f1d4439ce8b7093b32f99fa29b49621927eab149526bdc57d009c66b5b48907be8c31964e01b6d8ab2668e8470ed8f5004c57b306c116b5aa59798b82f6d91302c944922f44d0c24cec8b0563b68565bb5500b27b960e644b73c24ca527cdc691c211993f10d0040189fb7dc26207fcbeb8c078a44b4be67c42a1f6737027d0d5dc34dc1297a2eb3692330082404a02457ad3991b0372e56c4ee6d1af384e5b928da569ac58b4ab353f2ba72d4425ca19ca81e8b19bdee57b9364b29040962dad10aed4406508aebc45de9abef6c66d9c5d15aa7c6a7fc1727563a13d347e434ac51284f5c300d61aa0d6cc59d03f45ec20b36f8ff5e479463c5450ebc2a51c14c4ccaa6aafd38a81db42c37854d88fac41b489929d89fe0932925b2dd2bbd26161beec59eb25c4b3b3dcfc5d1caf3dcb77cef7b75d4b336f106642df20d37b1b417c919bb40c35a70061c17f8163aa505506c9fc7ad2277f01e62b01cb909fc6177bf3af26267ba4a049d4b9554d6bcd87315567be35b99f9b1f19ae8aa525df727ddddad6951279d5a65ef248920dd5cddf1316ec10980901084b5eecc2219e9dd1f872e7230fbad33b7e13091443f8a0b3874428eb3fe3b8225309cf2412477f2a71ceb7725ba431aaccdb098a64ceb63e190efda4123eb5a75d20dee27706f7b588650a0f73e9c8e1e8a000dd670fd2022b5caae6e4dc35a7cdc5417992eec88b0129db15a098376e7ee94279c48045c14b5b5679a210668f9450193ea94f5aefaae9112a3b654dae8118d9fe20cec16c6df3a25da6d5e7b2cbfe76c2c19d0fca355b44a7196718642fa1968de3f117175f5d7a8d83d5f929f4a93e9066fcb2517bd98bd29f54ab86238d73ae3932d8af9574725b531f02478a115604b5d4162bd659bf0a6ac81a5591c11b25c32cb2d6820d232bd9dada805967c079b830b0b890424aa9252583c5dcaa6b83dd16f2b22a5bd4d0b92014450e570c12c79272c83caffeb5c7d36c5f0e5f28c60a6f27a97784f95630e3d59b9d48f0d6037f2169cc833a4044da22c292ac80a4b89365449118c406a0d9e3cac55c60b0f8673eb360faf3339a0cf2dd2d0d386f032d9dc130cb310f172e983fa7cce4e98a400179bbd2e00f10610ee5feecc64bf577a37b3387269cb84e0af2ad3de39a933c392564800fad66e8cb1393f5aa4d545f807f0495178043715505a6706fe1c61ad4e8c047a874bf3577ca98687c9d6fdf08435cc077b91090ae2046e4a05f15cb61b5628f18c3d70d19d527ff4333cdfc3dd31189f1efe9f0b425c7847f1454b6bf7a5c5440722412f821baabf986411edc48045e863c5fa0f5be6ddaaf26273213fccceafb24fb5e0968849f52ccde60227d28f740a2ccfadb893033a060cc478ced8d77e5862e644fbc5b18e3dbb9e32041cf7757e3e2bcbefe048a4f1d4dee6093c1b305f9ce69a7f4a94798c2f9874369a5f9f449c6f67283b56d21865dd3b9d29613d1ce978d31a783fedc047269882e659d9b9133b95a520ba764234b52f2a90c31e2c964a3896cee3b855a03c98ef8ab64d017a3cc75abec1b0300b4169813552fe15e3e0674faa7c2c488cdc42ae95a5adf4f39f94178c0ee5f1848345aa0a4ba3817355ade":{
				Balance: 1000,
			},
		},
	}
}

func (s *AccountState) CreateAccount(address string) *Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc := &Account{Address: address}
	s.accounts[address] = acc
	return acc
}

func (s *AccountState) GetAccount(address string) (*Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.getAccountWithoutLock(address)
}

func (s *AccountState) getAccountWithoutLock(address string) (*Account, error) {
	account, ok := s.accounts[address]
	if !ok {
		return nil, ErrAccountNotFound
	}

	return account, nil
}

func (s *AccountState) GetBalance(address string) (uint64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	account, err := s.getAccountWithoutLock(address)
	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}

func (s *AccountState) Transfer(from, to string, amount uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fromAccount, err := s.getAccountWithoutLock(from)
	if err != nil {
		return err
	}

	// if fromAccount.Address.String() != "996fb92427ae41e4649b934ca495991b7852b855" {
	// 	if fromAccount.Balance < amount {
	// 		return ErrInsufficientBalance
	// 	}
	// }

	if fromAccount.Balance != 0 {
		fromAccount.Balance -= amount
	}

	if s.accounts[to] == nil {
		s.accounts[to] = &Account{
			Address: to,
		}
	}

	s.accounts[to].Balance += amount

	return nil
}