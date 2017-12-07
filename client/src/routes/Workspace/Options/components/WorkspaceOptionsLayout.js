/* @flow */

import React from 'react';
import Paper from 'material-ui/Paper';
import { withStyles } from 'material-ui/styles';
import type { SimulationOptions } from 'model/simulation/options';
import { FormInput, FormSelect, FormSwitch } from 'components/Form';
import { t } from 'i18n';

type Props = {
  options: SimulationOptions,
  updateOptionsField: (field: string, value: any) => void,
  classes: Object,
}

const scatteringTypeOptions = [
  { label: '', field: '' },
  { label: 'Gaussian', field: 'gaussian' },
  { label: 'Moliere', field: 'moliere' },
];

const energyStragglingOptions = [
  { label: '', field: '' },
  { label: 'Gaussian', field: 'gaussian' },
  { label: 'Vavilov', field: 'vavilov' },
];

class WorkspaceOptionsLayout extends React.Component<Props> {
  props: Props

  updateByType = (value: any, type: string) => {
    this.props.updateOptionsField(type, value);
  }

  render() {
    const { options, classes } = this.props;
    return (
      <div className={classes.root} >
        <Paper className={classes.item} >
          <FormSwitch
            type="antyparticleCorrectionOn"
            checked={options.antyparticleCorrectionOn}
            label={t('simulation.antyparticleCorrection')}
            onChange={this.updateByType}
          />
          <FormSwitch
            type="nuclearReactionsOn"
            checked={options.nuclearReactionsOn}
            label={t('simulation.nuclearReactions')}
            onChange={this.updateByType}
          />

          <FormInput
            type="meanEnergyLoss"
            value={options.meanEnergyLoss}
            label={t('simulation.meanEnergyLoss')}
            numbersOnly
            onChange={this.updateByType}
          />

          <FormInput
            type="minEnergyLoss"
            value={options.minEnergyLoss}
            label={t('simulation.minEnergyLoss')}
            numbersOnly
            onChange={this.updateByType}
          />
          <FormSelect
            type="scatteringType"
            value={options.scatteringType}
            label={t('simulation.scatteringType')}
            options={scatteringTypeOptions}
            onChange={this.updateByType}
          />
          <FormSelect
            type="energyStraggling"
            value={options.energyStraggling}
            label={t('simulation.energyStraggling')}
            options={energyStragglingOptions}
            onChange={this.updateByType}
          />

          <FormSwitch
            type="fastNeutronTransportOn"
            checked={options.fastNeutronTransportOn}
            label={t('simulation.fastNeutronTransport')}
            onChange={this.updateByType}
          />

          <FormInput
            type="lowEnergyNeutronCutOff"
            value={options.lowEnergyNeutronCutOff}
            label={t('simulation.lowEnergyNeutronCutOff')}
            numbersOnly
            onChange={this.updateByType}
          />
          <FormSwitch
            type="recordSecondaryNeutronCreation"
            checked={options.recordSecondaryNeutronCreation}
            label={t('workspace.editor.recordSecondaryNeutronCreation')}
            onChange={this.updateByType}
          />

          <FormInput
            type="numberOfGeneratedParticles"
            value={options.numberOfGeneratedParticles}
            label={t('workspace.editor.numberOfGeneratedParticles')}
            numbersOnly
            onChange={this.updateByType}
          />

          <FormInput
            type="numberOfRecordedParticles"
            value={options.numberOfRecordedParticles}
            label={t('workspace.editor.numberOfRecordedParticles')}
            numbersOnly
            onChange={this.updateByType}
          />
        </Paper>
      </div>
    );
  }
}

const styles = (theme: Object) => ({
  root: {
    margin: theme.spacing.unit,
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'flexStart',
  },
  item: {
    flex: '0 0 100%',
    margin: theme.spacing.unit,
    padding: theme.spacing.unit * 2,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'stretch',
  },
  stretched: {
    display: 'flex',
    flexDirection: 'row',
  },
});

export default withStyles(styles)(WorkspaceOptionsLayout);
