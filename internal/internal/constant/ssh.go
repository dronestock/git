package constant

const (
	SSHHome            = ".ssh"
	SSHConfigDir       = "config"
	SSHKeyFilename     = "id_rsa"
	SSHConfigFormatter = `Host *
  IgnoreUnknown UseKeychain
  UseKeychain yes
  AddKeysToAgent yes
  StrictHostKeyChecking=no
  IdentityFile %s
`
)
