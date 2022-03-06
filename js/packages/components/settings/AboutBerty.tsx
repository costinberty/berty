import React from 'react'
import { ScrollView, View } from 'react-native'

import { useStyles } from '@berty-tech/styles'
import { ScreenFC, useNavigation } from '@berty-tech/navigation'
import { useThemeColor } from '@berty-tech/store'

import { ButtonSettingV2, Section } from '../shared-components'

export const AboutBerty: ScreenFC<'Settings.AboutBerty'> = () => {
	const [{}, { scaleSize }] = useStyles()
	const colors = useThemeColor()
	const { navigate } = useNavigation()

	return (
		<View style={{ backgroundColor: colors['secondary-background'], flex: 1 }}>
			<ScrollView
				bounces={false}
				contentContainerStyle={{ paddingBottom: 12 * scaleSize }}
				showsVerticalScrollIndicator={false}
			>
				<Section>
					<ButtonSettingV2 text='FAQ' icon='bluetooth' onPress={() => navigate('Settings.Faq')} />
					<ButtonSettingV2 text='Roadmap' icon='info' />
					<ButtonSettingV2 text='Privacy Policy' icon='info' />
					<ButtonSettingV2 text='Open source licenses' icon='info' last />
				</Section>
			</ScrollView>
		</View>
	)
}
