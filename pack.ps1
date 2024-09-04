param (
    [string]$toolPath = ".\build\windows\amd64\ktor.exe",
    [string]$outPath = "ktor-installer.msi"
)

$version = $(git describe --tags --contains --always --abbrev=7)

if (!($version -match '\d+\.\d+.\d+')) {
    Write-Error "Expected version in the format *.*.*, got ${version}"
    exit 1
}

$wixProduct = @"
<Wix xmlns="http://wixtoolset.org/schemas/v4/wxs" xmlns:ui="http://wixtoolset.org/schemas/v4/wxs/ui">
    <Package Name="Ktor CLI" Version="${version}" Manufacturer="JetBrains" UpgradeCode="$(New-Guid)">
        <MediaTemplate EmbedCab="yes" />
        <WixVariable Id="WixUILicenseRtf" Value="LICENSE.rtf" />

        <MajorUpgrade
                AllowDowngrades="no" DowngradeErrorMessage="The newer version of Ktor CLI is already installed"
                AllowSameVersionUpgrades="yes"
        />

        <Property Id="WIXUI_EXITDIALOGOPTIONALTEXT" Value="Ktor CLI has been successfully installed. Use ktor.exe alias on the command line to launch the tool." />

        <StandardDirectory Id="ProgramFiles64Folder">
            <Directory Id="JetBrains" Name="JetBrains">
                <Directory Id="INSTALLDIR" Name="KtorCLI">
                    <Component Id="MainExecutable" Guid="$(New-Guid)">
                        <Environment Id="PATH" Name="PATH" Value="[INSTALLDIR]" Permanent="yes" Part="last" Action="set" System="no" />
                        <File Id="KtorExe" Name="ktor.exe" DiskId="1" Source="${toolPath}" KeyPath="yes" />
                    </Component>
                </Directory>
            </Directory>
        </StandardDirectory>

        <ui:WixUI Id="WixUI_InstallDir" InstallDirectory="INSTALLDIR" />
    </Package>
</Wix>
"@

$wixProduct | out-file -filepath KtorProduct.wxs
wix extension add -g WixToolset.UI.wixext
wix build -arch x64 -o $outPath -ext WixToolset.UI.wixext KtorProduct.wxs
Remove-Item KtorProduct.wxs