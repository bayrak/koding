JView             = require './../../core/jview'
mailServices      = [
  'gmail.com', 'hotmail.com', 'outlook.com', 'icloud.com'
  'yahoo.com', 'gmx.de', 'yandex.ru', 'yandex.com'
  'aol.com', 'mail.com', 'mail.ru'
  ]

module.exports = class TeamAllowedDomainTabForm extends KDFormView

  JView.mixin @prototype

  constructor:(options = {}, data)->

    options.cssClass = 'clearfix'

    super options, data

    team = KD.utils.getTeamData()
    if email = team.signup?.email
      domain = email.split('@').last
      domain = null  if domain in mailServices

    @label = new KDLabelView
      title : 'Allow sign up and team discovery with a specific email address'
      for   : 'allow'

    @checkbox = new KDInputView
      defaultValue : on
      type         : 'checkbox'
      name         : 'allow'
      label        : @label
      change       : =>
        state = @checkbox.getValue()
        if state
        then @showExtras()
        else @hideExtras()

    @input = new KDInputView
      placeholder  : 'domain.com, other.edu'
      name         : 'domains'
      defaultValue : domain  if domain

    @button = new KDButtonView
      title      : 'NEXT'
      style      : 'TeamsModal-button TeamsModal-button--green'
      attributes : testpath : 'allowed-domain-button'
      type       : 'submit'


  hideExtras: -> @$('div.extras').hide()

  showExtras: -> @$('div.extras').show()


  pistachio: ->

    """
    <div class='login-input-view tr'>{{> @checkbox}}{{> @label}}</div>
    <div class='extras'>
      <div class='login-input-view'><span>@</span>{{> @input}}</div>
      <p class='dim'>Please type a domain or multiple domains separated by commas.</p>
    </div>
    {{> @button}}
    """
